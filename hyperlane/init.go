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

package hyperlane

import (
	"log/slog"

	"jester.noble.xyz/metrics"
)

type Hyperlane struct {
	config  *config
	metrics *metrics.PrometheusMetrics
}

type config struct {
	hyperlaneMailbox      string
	hyperlaneNobleChainId uint32
}

type Overrides struct {
	HyperlaneMailbox      string
	HyperlaneNobleChainId uint32
}

func NewHyperlane(log *slog.Logger, testnet bool, metrics *metrics.PrometheusMetrics, overrides Overrides) *Hyperlane {
	return &Hyperlane{
		config:  newConfig(log, testnet, overrides),
		metrics: metrics,
	}
}

func newConfig(log *slog.Logger, testnet bool, overrides Overrides) *config {
	c := &config{}

	switch testnet {
	case true:
		c.hyperlaneMailbox = "0xfFAEF09B3cd11D9b20d1a19bECca54EEC2884766"
		c.hyperlaneNobleChainId = 0 // TODO
	default:
		c.hyperlaneMailbox = "0xc005dc82818d67AF737725bD4bf75435d065D239"
		c.hyperlaneNobleChainId = 0 // TODO
	}

	// Overrides
	if overrides.HyperlaneMailbox != "" {
		log.Info("overriding hyperlane mailbox contract address", "address", overrides.HyperlaneMailbox)
		c.hyperlaneMailbox = overrides.HyperlaneMailbox
	}
	if overrides.HyperlaneNobleChainId != 0 {
		log.Info("overriding hyperlane noble chain-id", "chainId", overrides.HyperlaneNobleChainId)
		c.hyperlaneNobleChainId = overrides.HyperlaneNobleChainId
	}

	return c
}
