// SPDX-License-Identifier: Apache-2.0
//
// Copyright 2025 NASD Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ethereum

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"jester.noble.xyz/metrics"
)

type Eth struct {
	config               *config
	metrics              *metrics.PrometheusMetrics
	WebsocketClient      *ethclient.Client
	websocketClientMutex sync.Mutex
	RPCClient            *ethclient.Client

	Redial *Redial
}

type config struct {
	websocketURL string
	rpcURL       string
}

// redial is used to manage the redial state between one gRPC client
// and multiple websockets.
type Redial struct {
	inProgressMutex sync.Mutex
	inProgress      bool
	cond            *sync.Cond
	err             error
	GetHistory      chan struct{} // redial done, trigger historical lookup
}

type clientType string

const (
	rpcClientType       clientType = "RPC"
	websocketClientType clientType = "Websocket"
)

// NewEth creates a new Ethereum instance with a websocket and rpc client.
// The intent behind this is to have this command run during cobras `PreRunE` or
// `PersistentPreRunE`.
// The returned *Eth pointer should be added to the app state.
func NewEth(
	ctx context.Context, log *slog.Logger, m *metrics.PrometheusMetrics,
	websocketURL, rpcURL string,
) (*Eth, error) {
	webSocketClient, err := dialClient(ctx, log, websocketURL, websocketClientType)
	if err != nil {
		return nil, fmt.Errorf("failed to dial websocket client: %v", err)
	}

	rpcClient, err := dialClient(ctx, log, rpcURL, rpcClientType)
	if err != nil {
		return nil, fmt.Errorf("failed to dial RPC client: %v", err)
	}

	config := &config{
		websocketURL: websocketURL,
		rpcURL:       rpcURL,
	}

	return &Eth{
		metrics:         m,
		config:          config,
		WebsocketClient: webSocketClient,
		RPCClient:       rpcClient,
		Redial:          newRedial(),
	}, nil
}

func newRedial() *Redial {
	return &Redial{
		GetHistory: make(chan struct{}),
	}
}

// dialClient creates an Ethereum Client.
// Based on the URL provided it will create either an RPC or Websocket client.
// The clientType parameter is used for logging purposes only.
func dialClient(ctx context.Context, log *slog.Logger, url string, clientType clientType) (client *ethclient.Client, err error) {
	err = retry.Do(func() error {
		client, err = ethclient.DialContext(ctx, url)
		if err != nil {
			return err
		}
		return nil
	},
		retry.Context(ctx),
		retry.OnRetry(func(attempt uint, err error) {
			log.Warn(fmt.Sprintf("retrying to dial Ethereum %s", clientType), "attempt", fmt.Sprintf("%d/%d", attempt+1, 10), "err", err)
		}))
	if err != nil {
		return nil, err
	}

	log.Info(fmt.Sprintf("successfully dialed Ethereum %s", clientType))

	return client, nil
}

// handleRedial handles the redial of the websocket client between multiple websocket subscriptions.
// Because the websocket client is shared between multiple subscriptions, this function
// is used to ensure that only one redial is in progress at a time.
func (e *Eth) handleRedial(ctx context.Context, log *slog.Logger) (err error) {
	redial := e.Redial
	redial.inProgressMutex.Lock()

	// Another goroutine is already handling redial
	if redial.inProgress {
		redial.inProgressMutex.Unlock()
		log.Info("client redial already in progress")
		redial.cond.L.Lock()   // Lock mutex to call wait()
		redial.cond.Wait()     // Wait for the redial to complete; unlocks mutex, waits for Broadcast(), re-locks mutex
		redial.cond.L.Unlock() // Unlock mutex
		errExists := false
		if redial.err != nil {
			errExists = true
		}
		log.Info("received client redial complete", "error-exists", errExists)
		return redial.err
	}

	// Mark redial as in progress and prepare a new signal channel
	redial.inProgress = true
	redial.cond = sync.NewCond(&redial.inProgressMutex)
	redial.inProgressMutex.Unlock()

	defer func() {
		redial.inProgressMutex.Lock()
		redial.inProgress = false
		redial.err = err
		redial.cond.Broadcast() // Signal all waiting goroutines that the redial operation is complete
		redial.inProgressMutex.Unlock()
	}()

	e.metrics.IncEthSubInterruptionCounter()

	client, err := dialClient(ctx, log, e.config.websocketURL, websocketClientType)
	if err != nil {
		return err
	}
	e.websocketClientMutex.Lock()
	e.WebsocketClient = client
	e.websocketClientMutex.Unlock()

	time.Sleep(1 * time.Second)     // Allow other redial signals to accumulate
	redial.GetHistory <- struct{}{} // Trigger historical lookup to catch up on missed events
	return nil
}

// filterLogs uses an RPC client to query Ethereum within a specified block range.
// It returns filtered logs based on contract address and topics.
func (e *Eth) FilterLogs(ctx context.Context, start, end *big.Int, contractAddress string, topics [][]common.Hash) ([]ethTypes.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: start,
		ToBlock:   end,
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		Topics:    topics,
	}
	return e.RPCClient.FilterLogs(ctx, query)
}

func (e *Eth) CloseClients() {
	if e.WebsocketClient != nil {
		e.WebsocketClient.Close()
	}
	if e.RPCClient != nil {
		e.RPCClient.Close()
	}
}
