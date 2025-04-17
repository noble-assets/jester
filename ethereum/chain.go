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
	"sync/atomic"
	"time"

	"github.com/avast/retry-go/v4"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"jester.noble.xyz/metrics"
)

type Eth struct {
	config               *config
	metrics              *metrics.PrometheusMetrics
	WebsocketClient      *ethclient.Client
	websocketClientMutex sync.Mutex
	RPCClient            *ethclient.Client

	currentHeight         atomic.Int64
	averageBlockTime      time.Duration
	averageBlockTimeMutex sync.Mutex

	ensureFinalityCh      chan *ethEventForFinality
	finalizedHeightPoller *finalizedHeightPoller

	Redial *Redial
}

type config struct {
	websocketURL string
	rpcURL       string
}

// redial is used to manage the redial state between one gRPC client
// and multiple websockets.
type Redial struct {
	inProgressMutex   sync.Mutex
	inProgress        bool
	LastObservedBlock int64
	cond              *sync.Cond
	err               error
	GetHistory        chan struct{} // redial done, trigger historical lookup
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

	e := Eth{
		metrics:               m,
		config:                config,
		WebsocketClient:       webSocketClient,
		RPCClient:             rpcClient,
		Redial:                newRedial(),
		ensureFinalityCh:      make(chan *ethEventForFinality, 10000),
		finalizedHeightPoller: newFinalizedHeightPoller(),
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	header, err := e.headerByNumber(ctxWithTimeout, log, big.NewInt(rpc.FinalizedBlockNumber.Int64()))
	if err != nil {
		return nil, fmt.Errorf("failed to get finalized block header: %v", err)
	}
	e.finalizedHeightPoller.currentFinalizedHeight.Store(header.Number.Uint64())

	return &e, nil
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

// CloseClients closes the websocket and RPC clients.
func (e *Eth) CloseClients() {
	if e.WebsocketClient != nil {
		e.WebsocketClient.Close()
	}
	if e.RPCClient != nil {
		e.RPCClient.Close()
	}
}

// trackCurrentHeight tracks the current Ethereum block height.
func (e *Eth) trackCurrentHeight(ctx context.Context, log *slog.Logger) error {
	return StartEventListener(
		ctx, log, e, "newHeads", "NewHead",
		func(ctx context.Context, sink chan *ethTypes.Header) (event.Subscription, error) {
			return e.WebsocketClient.SubscribeNewHead(ctx, sink)
		},
		func(ctx context.Context, log *slog.Logger, event *ethTypes.Header) {
			e.currentHeight.Store(event.Number.Int64())
		},
	)
}

// GetCurrentHeight returns the current Ethereum block height.
func (e *Eth) GetCurrentHeight() int64 {
	return e.currentHeight.Load()
}

const blockTimeWindow = 20

// trackAverageBlockTime calculates and tracks the average block time over the last blockTimeWindow blocks.
// It updates the average block time every hour.
func (e *Eth) trackAverageBlockTime(ctx context.Context, log *slog.Logger) error {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	blockDiff := big.NewInt(blockTimeWindow)

	calcAvg := func() error {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		latestHeader, err := e.headerByNumber(ctxWithTimeout, log, nil)
		if err != nil {
			return fmt.Errorf("failed to get latest block header when calculating average block time: %w", err)
		}

		oldBlockNum := new(big.Int).Sub(latestHeader.Number, blockDiff)

		oldHeader, err := e.headerByNumber(ctxWithTimeout, log, oldBlockNum)
		if err != nil {
			return fmt.Errorf("failed to get historical block header when calculating average block time: %w", err)
		}

		timeDiff := time.Duration(latestHeader.Time-oldHeader.Time) * time.Second
		averageBlockTime := timeDiff / time.Duration(blockTimeWindow)

		e.averageBlockTimeMutex.Lock()
		e.averageBlockTime = averageBlockTime
		e.averageBlockTimeMutex.Unlock()

		log.Info("updated average block time",
			"blockTime", averageBlockTime,
		)

		// TODO: metrics

		return nil
	}

	calcAvg()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := calcAvg(); err != nil {
				return fmt.Errorf("failed to calculate average block time: %w", err)
			}
		}
	}
}

func (e *Eth) GetAverageBlockTime() time.Duration {
	e.averageBlockTimeMutex.Lock()
	defer e.averageBlockTimeMutex.Unlock()
	return e.averageBlockTime
}
