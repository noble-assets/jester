package cmd

import (
	"fmt"
	"time"

	"github.com/noble-assets/jester/appstate"
	eth "github.com/noble-assets/jester/ethereum"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
func startCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "TODO",
		PreRunE: func(cmd *cobra.Command, _ []string) (err error) {
			a.Eth, err = eth.InitializeEth(a.Config.Ethereum.WebsocketURL, a.Config.Ethereum.RPCURL, cmd.Context())
			return
		},
		PostRun: func(_ *cobra.Command, _ []string) { a.Eth.CloseClients() },
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			log := a.Log
			ws := a.EthWebsocketClient
			g, _ := errgroup.WithContext(ctx) // TODO: should we use context returned here?

			// To get the data needed to query wormhole, we need to subscribe to two events:
			// 	1) MTokenSent
			// 	2) LogMessagePublished
			//
			// Log MessagePublished events are generic to the Wormhole, we only care about the
			// LogMessagePublished events that are also tied to a MTokenSent event.
			//
			// As we observe LogMessagePublished events, we map transaction hashes to the
			// sequence number emitted by the logMessagePublished event in the logMessagePublishedMap.
			// When we witness MTokenSent events, we query the map by the transaction hash
			// to get the relevant sequence number.
			//
			// We occasionally cleanup and delete LogMessagePublished events that do not have
			// a corresponding MTokenSent event.
			logMessagePublishedMap := eth.NewLogMessagePublishedMap()

			processingQueue := make(chan *eth.QueryData, 1000)

			g.Go(func() error {
				return eth.WormholeListener(ctx, logMessagePublishedMap, log, ws)
			})

			g.Go(func() error {
				return eth.M0Listener(ctx, logMessagePublishedMap, log, ws, processingQueue)
			})

			// Cleanup irrelevant LogMessagePublished events
			go func() {
				ticker := time.NewTicker(10 * time.Second)
				defer ticker.Stop()
				for range ticker.C {
					logMessagePublishedMap.Cleanup()
				}
			}()

			// Kick off workers to handle querying wormhole API to retrieve VAA's.
			// Due to rate limits with with the wormhole API, there is a fine balance between
			// workers and "GET" retries in `fetchVaa` function.
			maxWorkers := int64(3)
			sem := semaphore.NewWeighted(maxWorkers)
			g.Go(func() error {
				for {
					select {
					case <-ctx.Done():
						return nil
					case dequeued, ok := <-processingQueue:
						if !ok {
							panic("processingQueue channel closed")
						}
						if err := sem.Acquire(ctx, 1); err != nil {
							panic(fmt.Errorf("failed to acquire semaphore %w", err))
						}

						go func(data *eth.QueryData) {
							defer sem.Release(1)
							eth.StartQueryWorker(ctx, log, data)
						}(dequeued)
					}
				}
			})

			if err := g.Wait(); err != nil {
				return err
			}

			return nil
		},
	}

	appstate.AddConfigurationFlags(cmd)

	return cmd
}
