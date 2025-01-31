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
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"jester.noble.xyz/appstate"
	eth "jester.noble.xyz/ethereum"
	"jester.noble.xyz/metrics"
	"jester.noble.xyz/server"
	"jester.noble.xyz/state"
	"jester.noble.xyz/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
func startCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start Jester",
		Long: `Jester is a sidecar application designed to be run by the Noble validator set.

Jester facilitates the implementation of the Noble Dollar, powered by M0.

When started, Jester listens for events on the Ethereum blockchain and queries the Wormhole API for VAAs (Verifiable Action Approvals).
These VAAs are accumulated and served via a gRPC endpoint. This data acts as an ABCI Vote Extension.

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
				a.Config.Testnet,
				getOverrides(a.Viper),
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

			// Event subscription logic:
			// 	We watch for three Ethereum events:
			//  	1. `LogMessagePublished` from wormhole
			// 		2. `MTokenSent` from M0
			// 		3. `MTokenIndexSent` from M0
			//
			// As `LogMessagePublished` events are observed, their transaction hashes are mapped
			// to the sequence number emitted. This mapping is stored in `logMessagePublishedMap`.
			// This sequence number is needed later to query the Wormhole API.
			//
			// As we observe `MTokenSent` and `MTokenIndexSent` events, we use the tx hash to look up
			// the sequence number from the corresponding `LogMessagePublished` event. This event happens
			// in the same transaction as the M0 event. This metadata is then used to query the Wormhole API
			// for the VAA associated with the event.
			//
			// Note:
			//	`LogMessagePublished` events are generic to Wormhole and appear often.
			// 	While we do initially store all `LogMessagePublished` events, we periodically cleanup old entries.
			//
			// The `processingQueue` is a channel that holds `QueryData` structs. These structs
			// contain the necessary metadata to query the Wormhole API for VAAs.
			//
			// The `vaaList` is a state object that holds the VAAs we accumulate from the Wormhole API.
			// This list is served via gRPC to the Noble binary. Hitting this endpoint returns accumulated
			// VAA's and then clears the list. This is meant to only be queried by the Noble binary requesting the
			// Vote Extension.

			logMessagePublishedMap := state.NewLogMessagePublishedMap()
			processingQueue := make(chan *utils.QueryData, 1000)
			vaaList := state.NewVaaList()

			// Start Prometheus metrics server
			if a.Config.Metrics.Enabled {
				g.Go(func() error {
					return a.Metrics.StartServer(ctx, log, a.Mux, a.Config.Metrics.Address)
				})
			} else {
				log.Warn("prometheus metrics server disabled")
			}

			g.Go(func() error {
				return a.Eth.StartWormholeLMPListener(ctx, log, logMessagePublishedMap)
			})

			g.Go(func() error {
				return a.Eth.StartM0TokenSentListener(ctx, log, logMessagePublishedMap, processingQueue)
			})

			g.Go(func() error {
				return a.Eth.StartM0MTokenIndexSentListener(ctx, log, logMessagePublishedMap, processingQueue)
			})

			// Watch for websocket interruptions and get historical data to catch up.
			go func() {
				a.Eth.WatchForHistoryTrigger(ctx, log, processingQueue)
			}()

			// logMessagePublished events without an accompanying M0 event are irrelevant to us.
			// We periodically cleanup old entries from the map.
			go func() {
				ticker := time.NewTicker(10 * time.Second)
				defer ticker.Stop()
				for {
					select {
					case <-ctx.Done():
						return
					case <-ticker.C:
						logMessagePublishedMap.Cleanup()
					}
				}
			}()

			gRPCServer := server.NewJesterGrpcServer(log, a.Mux, a.Config.ServerAddress, vaaList)
			g.Go(func() error {
				return gRPCServer.StartServer(ctx)
			})

			// Worker pool to query the Wormhole API for VAAs.
			// Due to rate limits with wormhole's API, there is a fine balance between
			// `maxWorkers`` and "GET" retries in `fetchVaa`.
			maxWorkers := int64(3)
			sem := semaphore.NewWeighted(maxWorkers)
			g.Go(func() error {
				for {
					select {
					case <-ctx.Done():
						return nil
					case dequeued, ok := <-processingQueue:
						if !ok {
							return errors.New("processingQueue channel closed unexpectedly")
						}
						if err := sem.Acquire(ctx, 1); err != nil {
							return errors.Join(errors.New("failed to acquire semaphore"), err)
						}

						go func(data *utils.QueryData) {
							defer sem.Release(1)
							utils.StartWormholeWorker(
								ctx, log,
								a.Eth.Config.WormholeApiUrl,
								a.Eth.Config.WormholeTransceiver,
								a.Eth.Config.WormholeSrcChainId,
								data, vaaList,
							)
						}(dequeued)
					}
				}
			})

			// Get history on Jester start if startBlock is set
			startBlock := a.Viper.GetInt64(appstate.FlagStartBlock)
			if startBlock != 0 {
				endBlock := a.Viper.GetInt64(appstate.FlagEndBlock)
				go func() {
					a.Eth.GetHistory(ctx, log, processingQueue, startBlock, endBlock)
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

	// Contract and configuration override flags
	cmd.Flags().Uint16(appstate.FlagOverrideWormholeSrcChainId, 0, "override Wormhole source chain ID. This is likely the chain ID associated with Ethereum on Wormhole")
	cmd.Flags().Uint16(appstate.FlagOverrideNobleChainID, 0, "override noble Wormhole chain ID")
	cmd.Flags().String(appstate.FlagOverrideWormholeApiUrl, "", "override wormhole API URL")
	cmd.Flags().String(appstate.FlagOverrideHubPortal, "", "override the hub portal contract address")
	cmd.Flags().String(appstate.FlagOverrideWormholeCore, "", "override the wormhole core contract address")
	cmd.Flags().String(appstate.FlagOverrideWormholeTransceiver, "", "override the wormhole transceiver contract address")

	// misc flags
	cmd.Flags().BoolP(appstate.FlagHushLogo, "l", false, "suppress logo")

	return cmd
}

func getOverrides(v *viper.Viper) eth.Overrides {
	return eth.Overrides{
		WormholeSrcChainId:   v.GetUint16(appstate.FlagOverrideWormholeSrcChainId),
		WormholeNobleChainID: v.GetUint16(appstate.FlagOverrideNobleChainID),
		WormholeApiUrl:       v.GetString(appstate.FlagOverrideWormholeApiUrl),
		HubPortal:            v.GetString(appstate.FlagOverrideHubPortal),
		WormholeCore:         v.GetString(appstate.FlagOverrideWormholeCore),
		WormholeTransceiver:  v.GetString(appstate.FlagOverrideWormholeTransceiver),
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
