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

import "github.com/spf13/cobra"

type Config struct {
	Log_level     string   `toml:"log-level"`
	Log_style     string   `toml:"log-style"`
	Testnet       bool     `toml:"testnet"`
	ServerAddress string   `toml:"server-address"`
	Metrics       *Metrics `toml:"metrics"`

	Ethereum *Ethereum `toml:"ethereum"`
}

type Ethereum struct {
	WebsocketURL string `toml:"websocket-url"`
	RPCURL       string `toml:"rpc-url"`
}

type Metrics struct {
	Enabled bool   `toml:"enabled"`
	Address string `toml:"address"`
}

// AddConfigFlags adds flags that correspond to the config.toml file.
// The user has the option to use the flag, env var, or config file. Note that it
// takes precedent in that order.
//
// These are persistent flags.
// Ideally there is a flag for each config setting.
func AddConfigFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&eth_websocket, flagEthWebsocket, "w", "", "ethereum websocket")
	cmd.PersistentFlags().StringVarP(&eth_rpc, flagEthRPC, "r", "", "ethereum rpc")
	cmd.PersistentFlags().BoolVar(&testnet, flagTestnet, false, "use testnet configuration (contracts, chain ID's, ect.)")
	cmd.PersistentFlags().StringVar(&server_address, flagServerAddr, "localhost:9091", "gRPC server address")
	cmd.PersistentFlags().BoolVar(&metrics_enabled, flagMetricsEnabled, true, "enable prometheus metrics")
	cmd.PersistentFlags().StringVar(&metrics_address, flagMetricsAddress, "localhost:2112", "prometheus metrics address")
}
