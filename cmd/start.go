package cmd

import (
	"errors"
	"time"

	"github.com/noble-assets/jester/appstate"
	eth "github.com/noble-assets/jester/ethereum"
	"github.com/noble-assets/jester/server"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
func startCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start Jester",
		Long: `Jester is a "sidecar" application meant to be run by the Noble validator set.
		
Jester supports the implementation of the Noble Dollar, powered by M0.

Starting Jester will listen for events on the Ethereum blockchain and query the Wormhole API for VAAs.
Those VAAs are then accumulated and served via a gRPC endpoint.

Jesters default gRPC address/port is localhost:9091. This assumes that port 9090 is being used by the 
Noble binary. This can be overridden with the --server_address flag.

NOTE: The gRPC port for Jester is intended for use only by the Noble binary. 
Querying the gRPC endpoint "GetVaas" retrieves accumulated Wormhole VAAs and clears the state.

The Ethereum contracts are hard coded and can be toggled between mainnet and testnet via the --testnet flag.
Contracts and configurations can be overridden with the relevant "override" flags.
`,
		PreRunE: func(cmd *cobra.Command, _ []string) (err error) {
			a.Eth, err = eth.InitializeEth(
				cmd.Context(),
				a.Log,
				a.Config.Ethereum.WebsocketURL,
				a.Config.Ethereum.RPCURL,
				a.Config.Testnet,
				getEthOverrides(),
			)
			return
		},
		PostRun: func(_ *cobra.Command, _ []string) { a.Eth.CloseClients() },
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

			g.Go(func() error {
				return eth.WormholeListener(ctx, log, logMessagePublishedMap, a.Eth)
			})

			g.Go(func() error {
				return eth.M0Listener(ctx, log, logMessagePublishedMap, a.Eth, processingQueue)
			})

			// Watch for event subscriptions and get historical data
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

			// Start GRPC server
			vaaList := server.InitVaaList()
			go server.StartServer(ctx, log, a.Config.ServerAddress)

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
			startBlock := viper.GetInt64(appstate.FlagStartBlock)
			if startBlock != 0 {
				endBlock := viper.GetInt64(appstate.FlagEndBlock)
				go func() {
					eth.GetHistory(ctx, log, a.Eth, processingQueue, startBlock, endBlock)
				}()
			}

			if err := g.Wait(); err != nil {
				log.Error("fatal error", "error", err)
				return err
			}

			return nil
		},
	}

	appstate.AddConfigurationFlags(cmd)

	// Historical query flags
	cmd.Flags().Int64(appstate.FlagStartBlock, 0, "block number to start ethereum historical query, 0 for latest height")
	if err := viper.BindPFlag(appstate.FlagStartBlock, cmd.Flags().Lookup(appstate.FlagStartBlock)); err != nil {
		panic(err)
	}
	cmd.Flags().Int64(appstate.FlagEndBlock, 0, "block number to end ethereum historical query")
	if err := viper.BindPFlag(appstate.FlagEndBlock, cmd.Flags().Lookup(appstate.FlagEndBlock)); err != nil {
		panic(err)
	}

	// Contract and configuration override flags
	cmd.Flags().Uint16(appstate.FlagOverrideWormholeSrcChainId, 0, "override Wormhole source chain ID. This is likely the chain ID associated with Ethereum on Wormhole")
	if err := viper.BindPFlag(appstate.FlagOverrideWormholeSrcChainId, cmd.Flags().Lookup(appstate.FlagOverrideWormholeSrcChainId)); err != nil {
		panic(err)
	}
	cmd.Flags().Uint16(appstate.FlagOverrideNobleChainID, 0, "override noble Wormhole chain ID")
	if err := viper.BindPFlag(appstate.FlagOverrideNobleChainID, cmd.Flags().Lookup(appstate.FlagOverrideNobleChainID)); err != nil {
		panic(err)
	}
	cmd.Flags().String(appstate.FlagOverrideWormholeApiUrl, "", "override wormhole API URL")
	if err := viper.BindPFlag(appstate.FlagOverrideWormholeApiUrl, cmd.Flags().Lookup(appstate.FlagOverrideWormholeApiUrl)); err != nil {
		panic(err)
	}
	cmd.Flags().String(appstate.FlagOverrideHubPortal, "", "override the hub portal contract address")
	if err := viper.BindPFlag(appstate.FlagOverrideHubPortal, cmd.Flags().Lookup(appstate.FlagOverrideHubPortal)); err != nil {
		panic(err)
	}
	cmd.Flags().String(appstate.FlagOverrideWormholeCore, "", "override the wormhole core contract address")
	if err := viper.BindPFlag(appstate.FlagOverrideWormholeCore, cmd.Flags().Lookup(appstate.FlagOverrideWormholeCore)); err != nil {
		panic(err)
	}
	cmd.Flags().String(appstate.FlagOverrideWormholeTransceiver, "", "override the wormhole transceiver contract address")
	if err := viper.BindPFlag(appstate.FlagOverrideWormholeTransceiver, cmd.Flags().Lookup(appstate.FlagOverrideWormholeTransceiver)); err != nil {
		panic(err)
	}

	return cmd
}

func getEthOverrides() eth.Overrides {
	return eth.Overrides{
		WormholeSrcChainId:   viper.GetUint16(appstate.FlagOverrideWormholeSrcChainId),
		WormholeNobleChainID: viper.GetUint16(appstate.FlagOverrideNobleChainID),
		WormholeApiUrl:       viper.GetString(appstate.FlagOverrideWormholeApiUrl),
		HubPortal:            viper.GetString(appstate.FlagOverrideHubPortal),
		WormholeCore:         viper.GetString(appstate.FlagOverrideWormholeCore),
		WormholeTransceiver:  viper.GetString(appstate.FlagOverrideWormholeTransceiver),
	}
}
