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
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"jester.noble.xyz/metrics"
	"jester.noble.xyz/utils"
)

type Eth struct {
	Config               *Config
	Metrics              *metrics.PrometheusMetrics
	WebsocketClient      *ethclient.Client
	WebsocketClientMutex sync.Mutex
	RPCClient            *ethclient.Client

	redial *redial
}

type Config struct {
	WebsocketURL string
	RPCURL       string

	WormholeSrcChainId        uint16
	WormholeNobleChainID      uint16
	WormholeApiUrl            string
	HubPortal                 string
	WormholeCore              string
	WormholeTransceiver       string // LogMessagePublished topic sender
	PaddedWormholeTransceiver string
}

// redial is used to manage the redial state between one gRPC client
// and multiple websockets.
type redial struct {
	inProgressMutex sync.Mutex
	inProgress      bool
	cond            *sync.Cond
	err             error
	getHistory      chan struct{} // redial done, trigger historical lookup
}

type Overrides struct {
	WormholeSrcChainId   uint16
	WormholeNobleChainID uint16
	WormholeApiUrl       string
	HubPortal            string
	WormholeCore         string
	WormholeTransceiver  string
}

type clientType string

const (
	RPCClientType       clientType = "RPC"
	WebsocketClientType clientType = "Websocket"
)

// NewEth creates a new Ethereum instance with a websocket and rpc client.
// The intent behind this is to have this command run during cobras `PreRunE` or
// `PersistentPreRunE`.
// The returned *Eth pointer should be added to the app state.
func NewEth(
	ctx context.Context, log *slog.Logger, m *metrics.PrometheusMetrics,
	websocketurl, rpcurl string, testnet bool, overrides Overrides,
) (*Eth, error) {
	webSocketClient, err := dialClient(ctx, log, websocketurl, WebsocketClientType)
	if err != nil {
		return nil, fmt.Errorf("failed to dial websocket client: %v", err)
	}

	rpcClient, err := dialClient(ctx, log, rpcurl, RPCClientType)
	if err != nil {
		return nil, fmt.Errorf("failed to dial RPC client: %v", err)
	}

	return &Eth{
		Metrics:         m,
		Config:          newConfig(log, websocketurl, rpcurl, testnet, overrides),
		WebsocketClient: webSocketClient,
		RPCClient:       rpcClient,
		redial:          newRedial(),
	}, nil
}

// newConfig creates a new Ethereum Config
func newConfig(log *slog.Logger, websocketurl, rpcurl string, testnet bool, overrides Overrides) *Config {
	c := &Config{
		WebsocketURL: websocketurl,
		RPCURL:       rpcurl,

		WormholeNobleChainID: 4009,
	}

	switch testnet {
	case true:
		c.WormholeSrcChainId = 10002
		c.WormholeApiUrl = "https://api.testnet.wormholescan.io/v1/signed_vaa"
		// https://github.com/m0-foundation/m-portal/blob/682481178808005a160e41d5318242c1abc2f88f/deployments/11155111.json#L3
		c.HubPortal = "0xD925C84b55E4e44a53749fF5F2a5A13F63D128fd"
		// https://wormhole.com/docs/build/reference/contract-addresses/#__tabbed_1_2
		c.WormholeCore = "0x4a8bc80Ed5a4067f1CCf107057b8270E0cC11A78"
		// https://github.com/m0-foundation/m-portal/blob/682481178808005a160e41d5318242c1abc2f88f/deployments/11155111.json#L5
		c.WormholeTransceiver = "0x0763196A091575adF99e2306E5e90E0Be5154841"
	default:
		c.WormholeSrcChainId = 2
		c.WormholeApiUrl = "https://api.wormholescan.io/v1/signed_vaa"
		// https://github.com/m0-foundation/m-portal/blob/dbe93da561c94dfc04beec8a144b11b287957b7a/deployments/noble/1.json#L2
		c.HubPortal = "0x83Ae82Bd4054e815fB7B189C39D9CE670369ea16"
		// https://wormhole.com/docs/build/reference/contract-addresses/#__tabbed_1_1
		c.WormholeCore = "0x98f3c9e6E3fAce36bAAd05FE09d375Ef1464288B"
		// https://github.com/m0-foundation/m-portal/blob/dbe93da561c94dfc04beec8a144b11b287957b7a/deployments/noble/1.json#L3
		c.WormholeTransceiver = "0xc7Dd372c39E38BF11451ab4A8427B4Ae38ceF644"
	}

	// Overrides
	if overrides.WormholeSrcChainId != 0 {
		log.Info("overriding wormhole source chain ID", "chainID", overrides.WormholeSrcChainId)
		c.WormholeSrcChainId = overrides.WormholeSrcChainId
	}
	if overrides.WormholeNobleChainID != 0 {
		log.Info("overriding noble wormhole chain ID", "chainID", overrides.WormholeNobleChainID)
		c.WormholeNobleChainID = overrides.WormholeNobleChainID
	}
	if overrides.WormholeApiUrl != "" {
		log.Info("overriding wormhole API URL", "url", overrides.WormholeApiUrl)
		c.WormholeApiUrl = overrides.WormholeApiUrl
	}
	if overrides.HubPortal != "" {
		log.Info("overriding hub portal contract", "address", overrides.HubPortal)
		c.HubPortal = overrides.HubPortal
	}
	if overrides.WormholeCore != "" {
		log.Info("overriding wormhole core contract", "address", overrides.WormholeCore)
		c.WormholeCore = overrides.WormholeCore
	}
	if overrides.WormholeTransceiver != "" {
		log.Info("overriding wormhole transceiver contract", "address", overrides.WormholeTransceiver)
		c.WormholeTransceiver = overrides.WormholeTransceiver
	}

	paddedWormholeTransceiver := make([]byte, 32)
	copy(paddedWormholeTransceiver[12:], common.FromHex(c.WormholeTransceiver))
	c.PaddedWormholeTransceiver = common.Bytes2Hex(paddedWormholeTransceiver)

	return c
}

func newRedial() *redial {
	return &redial{
		getHistory: make(chan struct{}),
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
	redial := e.redial
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

	e.Metrics.IncEthSubInterruptionCounter()

	client, err := dialClient(ctx, log, e.Config.WebsocketURL, WebsocketClientType)
	if err != nil {
		return err
	}
	e.WebsocketClientMutex.Lock()
	e.WebsocketClient = client
	e.WebsocketClientMutex.Unlock()

	time.Sleep(1 * time.Second)     // Allow other redial signals to accumulate
	redial.getHistory <- struct{}{} // Trigger historical lookup to catch up on missed events
	return nil
}

// WatchForHistoryTrigger is used to catch up on any block data missed during an event
// subscription interruption. It is hardcoded to look back 50 blocks.
//
// It is meant to be run in a goroutine
func (e *Eth) WatchForHistoryTrigger(ctx context.Context, log *slog.Logger, processingQueue chan *utils.QueryData) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-e.redial.getHistory:
			lookback := uint64(50) // TODO: make lookback configurable
			log.Info(fmt.Sprintf("getting historical events for %d blocks", lookback))
			latest, err := e.RPCClient.BlockNumber(ctx)
			if err != nil {
				log.Error("failed to get latest block number", "error", err)
				continue
			}
			start := latest - lookback
			e.GetHistory(ctx, log, processingQueue, int64(start), 0)
		}
	}
}

func (e *Eth) CloseClients() {
	if e.WebsocketClient != nil {
		e.WebsocketClient.Close()
	}
	if e.RPCClient != nil {
		e.RPCClient.Close()
	}
}
