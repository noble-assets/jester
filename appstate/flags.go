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

package appstate

const (
	FlagHome                        = "home"
	FlagLogLevel                    = "log-level"
	FlagLogStyle                    = "log-style"
	flagTestnet                     = "testnet"
	flagEthWebsocket                = "ethereum.websocket-url"
	flagEthRPC                      = "ethereum.rpc-url"
	flagServerAddr                  = "server-address"
	FlagStartBlock                  = "start-block"
	FlagEndBlock                    = "end-block"
	FlagOverrideWormholeSrcChainId  = "override-wormhole-src-chain-id"
	FlagOverrideNobleChainID        = "override-noble-chain-id"
	FlagOverrideWormholeApiUrl      = "override-wormhole-api-url"
	FlagOverrideHubPortal           = "override-hub-portal"
	FlagOverrideWormholeTransceiver = "override-wormhole-transceiver"
	FlagOverrideWormholeCore        = "override-wormhole-core"
	FlagOverrideFetchVAAAttempts    = "override-fetch-vaa-attempts"
	flagMetricsEnabled              = "metrics.enabled"
	flagMetricsAddress              = "metrics.address"
	FlagHushLogo                    = "hush-logo"
	FlagDeveloperMode               = "dev"
)

// if the flag is being added to multiple commands, you must use the flags that accepts pointers
// example `StringVar` instead of `String`
var (
	eth_websocket   string
	eth_rpc         string
	server_address  string
	testnet         bool
	metrics_enabled bool
	metrics_address string
)
