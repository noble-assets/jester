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

NOTE: The gRPC port for Jester is intended for use only by the Noble binary. 
Querying the gRPC endpoint "GetVaas" retrieves accumulated Wormhole VAAs and clears the state.

The Ethereum contracts are hard coded and can be toggled between mainnet and testnet via the --testnet flag.
The contracts can be overridden with the --override_mportal_contract, --override_wormhole_contract,
and --override_lmp_sender flags.
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
			ws := a.EthWebsocketClient
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
				return eth.WormholeListener(ctx, log, logMessagePublishedMap, ws, a.Eth.Config.WormholeContract, a.Eth.Config.LogMessagePublishedSender)
			})

			g.Go(func() error {
				return eth.M0Listener(ctx, log, logMessagePublishedMap, ws, processingQueue, a.Eth.Config.MPortalContract, a.Eth.Config.LogMessagePublishedSender)
			})

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
							eth.StartQueryWorker(ctx, log, data, vaaList)
						}(dequeued)
					}
				}
			})

			// Historical query
			startBlock := viper.GetInt64(appstate.FlagStartBlock)
			if startBlock != 0 {
				endBlock := viper.GetInt64(appstate.FlagEndBlock)
				go func() {
					eth.GetHistory(
						ctx, log,
						a.EthRPCClient,
						processingQueue,
						startBlock, endBlock,
						a.Eth.Config.MPortalContract,
						a.Eth.Config.WormholeContract,
						a.Eth.Config.LogMessagePublishedSender,
					)
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

	// Contract override flags
	cmd.Flags().String(appstate.FlagOverrideWormholeContract, "", "override wormhole contract address")
	if err := viper.BindPFlag(appstate.FlagOverrideWormholeContract, cmd.Flags().Lookup(appstate.FlagOverrideWormholeContract)); err != nil {
		panic(err)
	}
	cmd.Flags().String(appstate.FlagOverrideMPortalContract, "", "override M0's MPortal contract address")
	if err := viper.BindPFlag(appstate.FlagOverrideMPortalContract, cmd.Flags().Lookup(appstate.FlagOverrideMPortalContract)); err != nil {
		panic(err)
	}
	cmd.Flags().String(appstate.FlagOverrideLMPSender, "", "override LogMessagePublished event sender")
	if err := viper.BindPFlag(appstate.FlagOverrideLMPSender, cmd.Flags().Lookup(appstate.FlagOverrideLMPSender)); err != nil {
		panic(err)
	}

	return cmd
}

func getEthOverrides() eth.Overrides {
	return eth.Overrides{
		MPortalContract:           viper.GetString(appstate.FlagOverrideMPortalContract),
		WormholeContract:          viper.GetString(appstate.FlagOverrideWormholeContract),
		LogMessagePublishedSender: viper.GetString(appstate.FlagOverrideLMPSender),
	}
}
