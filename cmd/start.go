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
)

// startCmd represents the start command
func startCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start Jester",
		Long: `Jester is a "sidecar" application meant to be run by the Noble validator set.
		
Jester supports the implementation of the Noble Dollar, powered by M0.

NOTE: The gRPC port for Jester is intended for use only by the Noble binary. 
Querying the gRPC endpoint "GetVaas" retrieves accumulated Wormhole VAAs and clears the state.`,
		PreRunE: func(cmd *cobra.Command, _ []string) (err error) {
			a.Eth, err = eth.InitializeEth(cmd.Context(), a.Log, a.Config.Ethereum.WebsocketURL, a.Config.Ethereum.RPCURL)
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
				return eth.WormholeListener(ctx, log, logMessagePublishedMap, ws)
			})

			g.Go(func() error {
				return eth.M0Listener(ctx, log, logMessagePublishedMap, ws, processingQueue)
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
			go server.StartServer(ctx, log)

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

			if err := g.Wait(); err != nil {
				log.Error("fatal error", "error", err)
				return err
			}

			return nil
		},
	}

	appstate.AddConfigurationFlags(cmd)

	return cmd
}
