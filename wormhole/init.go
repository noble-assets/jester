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

package wormhole

import (
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"jester.noble.xyz/v2/metrics"
	"jester.noble.xyz/v2/wormhole/state"
)

type Wormhole struct {
	config                 *config
	metrics                *metrics.PrometheusMetrics
	logMessagePublishedMap *state.LogMessagePublishedMap
	processingQueue        chan *queryData
	VaaList                *state.VaaList
}

type config struct {
	wormholeSrcChainId        uint16
	wormholeNobleChainID      uint16
	wormholeApiUrl            string
	hubPortal                 string
	wormholeCore              string
	wormholeTransceiver       string // LogMessagePublished topic sender
	paddedWormholeTransceiver string

	fetchVAAAttempts uint
}

type Overrides struct {
	WormholeSrcChainId   uint16
	WormholeNobleChainID uint16
	WormholeApiUrl       string
	HubPortal            string
	WormholeCore         string
	WormholeTransceiver  string

	FetchVAAAttempts uint
}

func NewWormhole(log *slog.Logger, testnet bool, metrics *metrics.PrometheusMetrics, overrides Overrides) *Wormhole {
	return &Wormhole{
		config:                 newConfig(log, testnet, overrides),
		metrics:                metrics,
		logMessagePublishedMap: state.NewLogMessagePublishedMap(),
		processingQueue:        make(chan *queryData, 1000),
		VaaList:                state.NewVaaList(),
	}
}

func newConfig(log *slog.Logger, testnet bool, overrides Overrides) *config {
	c := &config{
		wormholeNobleChainID: 4009,
		fetchVAAAttempts:     overrides.FetchVAAAttempts,
	}

	switch testnet {
	case true:
		c.wormholeSrcChainId = 10002
		c.wormholeApiUrl = "https://api.testnet.wormholescan.io/v1/signed_vaa"
		// https://github.com/m0-foundation/m-portal/blob/682481178808005a160e41d5318242c1abc2f88f/deployments/11155111.json#L3
		c.hubPortal = "0xD925C84b55E4e44a53749fF5F2a5A13F63D128fd"
		// https://wormhole.com/docs/build/reference/contract-addresses/#__tabbed_1_2
		c.wormholeCore = "0x4a8bc80Ed5a4067f1CCf107057b8270E0cC11A78"
		// https://github.com/m0-foundation/m-portal/blob/682481178808005a160e41d5318242c1abc2f88f/deployments/11155111.json#L5
		c.wormholeTransceiver = "0x0763196A091575adF99e2306E5e90E0Be5154841"
	default:
		c.wormholeSrcChainId = 2
		c.wormholeApiUrl = "https://api.wormholescan.io/v1/signed_vaa"
		// https://github.com/m0-foundation/m-portal/blob/dbe93da561c94dfc04beec8a144b11b287957b7a/deployments/noble/1.json#L2
		c.hubPortal = "0x83Ae82Bd4054e815fB7B189C39D9CE670369ea16"
		// https://wormhole.com/docs/build/reference/contract-addresses/#__tabbed_1_1
		c.wormholeCore = "0x98f3c9e6E3fAce36bAAd05FE09d375Ef1464288B"
		// https://github.com/m0-foundation/m-portal/blob/dbe93da561c94dfc04beec8a144b11b287957b7a/deployments/noble/1.json#L3
		c.wormholeTransceiver = "0xc7Dd372c39E38BF11451ab4A8427B4Ae38ceF644"
	}

	// Overrides
	if overrides.WormholeSrcChainId != 0 {
		log.Info("overriding wormhole source chain ID", "chainID", overrides.WormholeSrcChainId)
		c.wormholeSrcChainId = overrides.WormholeSrcChainId
	}
	if overrides.WormholeNobleChainID != 0 {
		log.Info("overriding noble wormhole chain ID", "chainID", overrides.WormholeNobleChainID)
		c.wormholeNobleChainID = overrides.WormholeNobleChainID
	}
	if overrides.WormholeApiUrl != "" {
		log.Info("overriding wormhole API URL", "url", overrides.WormholeApiUrl)
		c.wormholeApiUrl = overrides.WormholeApiUrl
	}
	if overrides.HubPortal != "" {
		log.Info("overriding hub portal contract", "address", overrides.HubPortal)
		c.hubPortal = overrides.HubPortal
	}
	if overrides.WormholeCore != "" {
		log.Info("overriding wormhole core contract", "address", overrides.WormholeCore)
		c.wormholeCore = overrides.WormholeCore
	}
	if overrides.WormholeTransceiver != "" {
		log.Info("overriding wormhole transceiver contract", "address", overrides.WormholeTransceiver)
		c.wormholeTransceiver = overrides.WormholeTransceiver
	}

	paddedWormholeTransceiver := make([]byte, 32)
	copy(paddedWormholeTransceiver[12:], common.FromHex(c.wormholeTransceiver))
	c.paddedWormholeTransceiver = common.Bytes2Hex(paddedWormholeTransceiver)

	return c
}
