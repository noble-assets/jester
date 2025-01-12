package cmd

import (
	"errors"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"jester.noble.xyz/appstate"
	eth "jester.noble.xyz/ethereum"
	"jester.noble.xyz/noble"
	"jester.noble.xyz/server"
	"jester.noble.xyz/state"

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
These VAAs are accumulated and served via a gRPC endpoint.

By default, Jester's gRPC server listens on localhost:9091. This assumes that port 9090 is being used by the Noble binary. 
You can override this default address with the --server_address flag.

Note: The gRPC port for Jester is intended for use only by the Noble binary. 
Querying the gRPC endpoint "GetVoteExtention" retrieves accumulated Wormhole VAAs and clears the state.

The Ethereum contracts are hardcoded and can be toggled between mainnet and testnet using the --testnet flag.
You can override contracts and configurations with the relevant "override" flags.`,
		PreRunE: func(cmd *cobra.Command, _ []string) (err error) {
			a.Eth, err = eth.InitializeEth(
				cmd.Context(),
				a.Log,
				a.Config.Ethereum.WebsocketURL,
				a.Config.Ethereum.RPCURL,
				a.Config.Testnet,
				getEthOverrides(a.Viper),
			)
			if err != nil {
				return err
			}
			a.Noble, err = noble.InitializeNoble(cmd.Context(), a.Log, a.Config.Noble.GRPCURL, a.Viper.GetBool(noble.FlagSkipHealth))
			return err
		},
		PostRun: func(_ *cobra.Command, _ []string) {
			a.Eth.CloseClients()
			a.Noble.CloseClients()
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			log := a.Log
			g, ctx := errgroup.WithContext(ctx)

			// Event Subscription:
			// To query Wormhole data, we need to subscribe to two events:
			// 1. MTokenSent
			// 2. LogMessagePublished
			//
			// LogMessagePublished events are generic to Wormhole; we only process ones
			// tied to an MTokenSent event.
			//
			// As LogMessagePublished events are observed, their transaction hashes are mapped
			// to the sequence number emitted. This mapping is stored in `logMessagePublishedMap`.
			// When MTokenSent events occur, the map is queried using the transaction hash
			// to retrieve the associated sequence number.
			//
			// Irrelevant LogMessagePublished events without corresponding MTokenSent events
			// are periodically cleaned up.

			logMessagePublishedMap := eth.NewLogMessagePublishedMap()
			processingQueue := make(chan *eth.QueryData, 1000)
			vaaList := state.NewVaaList()

			g.Go(func() error {
				return eth.WormholeListener(ctx, log, logMessagePublishedMap, a.Eth)
			})

			g.Go(func() error {
				return eth.M0Listener(ctx, log, logMessagePublishedMap, a.Eth, processingQueue)
			})

			// Watch for event subscription interruptions and get historical data
			go func() {
				a.Eth.GetHistoricalOnRedial(ctx, log, processingQueue)
			}()

			// Cleanup irrelevant LogMessagePublished events
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

			gRPCServer := server.NewJesterGrpcServer(a.Config.ServerAddress, log, vaaList, a.Noble.WormholeClient)
			go gRPCServer.Start(ctx)

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

						go func(data *eth.QueryData) {
							defer sem.Release(1)
							eth.StartQueryWorker(ctx, log, a.Eth.Config.WormholeApiUrl, data, vaaList)
						}(dequeued)
					}
				}
			})

			// Historical query
			startBlock := a.Viper.GetInt64(appstate.FlagStartBlock)
			if startBlock != 0 {
				endBlock := a.Viper.GetInt64(appstate.FlagEndBlock)
				go func() {
					eth.GetHistory(ctx, log, a.Eth, processingQueue, startBlock, endBlock)
				}()
			}
			//
			if err := g.Wait(); err != nil {
				log.Error("fatal error", "error", err)
				return err
			}

			return nil
		},
	}

	appstate.AddConfigurationFlags(cmd)

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

	// testing flags
	cmd.Flags().Bool(noble.FlagSkipHealth, false, "skip Noble gRPC health check")
	cmd.Flags().Lookup(noble.FlagSkipHealth).Hidden = true

	return cmd
}

func getEthOverrides(v *viper.Viper) eth.Overrides {
	return eth.Overrides{
		WormholeSrcChainId:   v.GetUint16(appstate.FlagOverrideWormholeSrcChainId),
		WormholeNobleChainID: v.GetUint16(appstate.FlagOverrideNobleChainID),
		WormholeApiUrl:       v.GetString(appstate.FlagOverrideWormholeApiUrl),
		HubPortal:            v.GetString(appstate.FlagOverrideHubPortal),
		WormholeCore:         v.GetString(appstate.FlagOverrideWormholeCore),
		WormholeTransceiver:  v.GetString(appstate.FlagOverrideWormholeTransceiver),
	}
}
