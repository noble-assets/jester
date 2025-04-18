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

package cmd

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"jester.noble.xyz/appstate"
	eth "jester.noble.xyz/ethereum"
	"jester.noble.xyz/hyperlane"
	"jester.noble.xyz/metrics"
	"jester.noble.xyz/server"
	"jester.noble.xyz/wormhole"
)

// startCmd represents the start command
func startCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start Jester",
		Long: `Jester is a sidecar application designed to be run by the Noble validator set.

Jester facilitates the implementation of the Noble Dollar, powered by M0.

By default, Jester's gRPC server listens and serves on localhost:9091. This assumes that port 9090 is being used by the Noble binary. 
You can override this default address with the --server_address flag.

Note: The gRPC port for Jester is intended to be queried only by the Noble binary. 
Querying the gRPC endpoint "GetVoteExtension" retrieves accumulated Wormhole VAAs and then CLEARS the memory state.

The Ethereum contracts are hardcoded and can be toggled between mainnet and testnet using the --testnet flag.
You can override contracts and configurations with the relevant "override" flags.`,
		PreRunE: func(cmd *cobra.Command, _ []string) (err error) {
			if !a.Viper.GetBool(appstate.FlagHushLogo) {
				plogo()
			}
			a.Mux = http.NewServeMux()
			a.Metrics = metrics.NewPrometheusMetrics()
			a.Eth, err = eth.NewEth(
				cmd.Context(), a.Log, a.Metrics,
				a.Config.Ethereum.WebsocketURL,
				a.Config.Ethereum.RPCURL,
			)
			return err
		},
		PostRun: func(_ *cobra.Command, _ []string) {
			a.Eth.CloseClients()
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			log := a.Log
			g, ctx := errgroup.WithContext(ctx)
			wormholeOverrides, hyperlaneOverrides := getOverrides(a.Viper)

			if a.Config.Metrics.Enabled {
				g.Go(func() error {
					return a.Metrics.StartServer(ctx, log, a.Mux, a.Config.Metrics.Address)
				})
			} else {
				log.Warn("prometheus metrics server disabled")
			}

			g.Go(func() error {
				return a.Eth.Start(ctx, log)
			})

			w := wormhole.NewWormhole(log, a.Config.Testnet, a.Metrics, wormholeOverrides)
			g.Go(func() error {
				return w.Start(ctx, log, a.Eth, a.Metrics)
			})

			h := hyperlane.NewHyperlane(log, a.Config.Testnet, a.Metrics, hyperlaneOverrides)
			g.Go(func() error {
				return h.Start(ctx, log, a.Eth, a.Metrics)
			})

			// Watch for websocket interruptions and get historical data to catch up.
			// Since historical data is not crucial to Jester, we do not return an error.
			go func() {
				for {
					select {
					case <-ctx.Done():
						return
					case <-a.Eth.Redial.GetHistory:
						// out of precaution, look back an extra 2 minutes worth of blocks
						twoMinOfBlocks := (time.Duration(2) * time.Minute) / a.Eth.GetAverageBlockTime()
						lookBackStart := a.Eth.Redial.LastObservedBlock - int64(twoMinOfBlocks)

						log.Info(fmt.Sprintf("getting historical events for %d blocks", (a.Eth.GetCurrentHeight() - lookBackStart)))

						go func() { w.GetHistory(ctx, log, a.Eth, int64(lookBackStart), 0) }()
						go func() { h.GetHistory(ctx, log, a.Eth, int64(lookBackStart), 0) }()
					}
				}
			}()

			gRPCServer := server.NewJesterGrpcServer(log, a.Metrics, a.Mux, a.Config.ServerAddress, w.VaaList)
			g.Go(func() error {
				return gRPCServer.StartServer(ctx)
			})

			// Fetch historical data if startBlock is set.
			startBlock := a.Viper.GetInt64(appstate.FlagStartBlock)
			if startBlock != 0 {
				endBlock := a.Viper.GetInt64(appstate.FlagEndBlock)
				go func() { w.GetHistory(ctx, log, a.Eth, startBlock, endBlock) }()
				go func() { h.GetHistory(ctx, log, a.Eth, startBlock, endBlock) }()
			}

			// Developer mode
			if a.Viper.GetBool(appstate.FlagDeveloperMode) {
				log.Warn("developer mode enabled")
				go func() {
					ticker := time.NewTicker(1 * time.Hour)
					defer ticker.Stop()
					for {
						select {
						case <-ctx.Done():
							return
						case <-ticker.C:
							vaas := w.VaaList.GetThenClearAll()
							log.Info("developer mode: VAA's cleared", "vaa-count", len(vaas), "vaa-list", vaas)
						}
					}
				}()
			}

			// Hold here, waiting for any errors from errgroup
			if err := g.Wait(); err != nil {
				log.Error("fatal error", "error", err)
				return err
			}

			return nil
		},
	}

	appstate.AddConfigFlags(cmd)

	// Historical query flags
	cmd.Flags().Int64(appstate.FlagStartBlock, 0, "block number to start ethereum historical query")
	cmd.Flags().Int64(appstate.FlagEndBlock, 0, "block number to end ethereum historical query, 0 for latest height")

	// Wormhole overrides
	cmd.Flags().Uint16(appstate.FlagOverrideWormholeSrcChainId, 0, "override Wormhole source chain ID. This is likely the chain ID associated with Ethereum on Wormhole")
	cmd.Flags().Uint16(appstate.FlagOverrideNobleChainID, 0, "override noble Wormhole chain ID")
	cmd.Flags().String(appstate.FlagOverrideWormholeApiUrl, "", "override wormhole API URL")
	cmd.Flags().String(appstate.FlagOverrideHubPortal, "", "override the hub portal contract address")
	cmd.Flags().String(appstate.FlagOverrideWormholeCore, "", "override the wormhole core contract address")
	cmd.Flags().String(appstate.FlagOverrideWormholeTransceiver, "", "override the wormhole transceiver contract address")
	cmd.Flags().Uint(appstate.FlagOverrideFetchVAAAttempts, 60, "override the maximum number of attempts made to query the Wormhole API for VAAs (Attempts are spaced 30 seconds apart)")

	// Hyperlane overrides
	cmd.Flags().String(appstate.FlagOverrideHyperlaneMailbox, "", "override the hyperlane mailbox contract address")
	cmd.Flags().Uint32(appstate.FlagOverrideHyperlaneNobleChainId, 0, "override the hyperlane chain ID")

	// misc flags
	cmd.Flags().BoolP(appstate.FlagHushLogo, "l", false, "suppress logo")
	cmd.Flags().Bool(appstate.FlagDeveloperMode, false, "enable developer mode (clears VAA cache every hour; useful for non-validators tracking Prometheus metrics)")
	return cmd
}

func getOverrides(v *viper.Viper) (wormhole.Overrides, hyperlane.Overrides) {
	return wormhole.Overrides{
			WormholeSrcChainId:   v.GetUint16(appstate.FlagOverrideWormholeSrcChainId),
			WormholeNobleChainID: v.GetUint16(appstate.FlagOverrideNobleChainID),
			WormholeApiUrl:       v.GetString(appstate.FlagOverrideWormholeApiUrl),
			HubPortal:            v.GetString(appstate.FlagOverrideHubPortal),
			WormholeCore:         v.GetString(appstate.FlagOverrideWormholeCore),
			WormholeTransceiver:  v.GetString(appstate.FlagOverrideWormholeTransceiver),
			FetchVAAAttempts:     v.GetUint(appstate.FlagOverrideFetchVAAAttempts),
		},
		hyperlane.Overrides{
			HyperlaneMailbox:      v.GetString(appstate.FlagOverrideHyperlaneMailbox),
			HyperlaneNobleChainId: v.GetUint32(appstate.FlagOverrideHyperlaneNobleChainId),
		}
}

func plogo() {
	x := fmt.Sprintf(`
     ██╗███████╗███████╗████████╗███████╗██████╗ 
     ██║██╔════╝██╔════╝╚══██╔══╝██╔════╝██╔══██╗
     ██║█████╗  ███████╗   ██║   █████╗  ██████╔╝
██   ██║██╔══╝  ╚════██║   ██║   ██╔══╝  ██╔══██╗
╚█████╔╝███████╗███████║   ██║   ███████╗██║  ██║
 ╚════╝ ╚══════╝╚══════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝

                                 Noble's sidekick
                                 %s
`, Version)
	lines := strings.Split(x, "\n")
	for _, line := range lines {
		println(line)
		time.Sleep(100 * time.Millisecond)
	}
}
