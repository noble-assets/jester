package cmd

import (
	"fmt"
	"time"

	"github.com/noble-assets/jester/appstate"
	eth "github.com/noble-assets/jester/ethereum"
	"golang.org/x/sync/errgroup"

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
			g, _ := errgroup.WithContext(ctx)

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

			processingQueue := make(chan *eth.QueryData)

			g.Go(func() error {
				return eth.WormholeListener(ctx, logMessagePublishedMap, log, ws)
			})

			g.Go(func() error {
				return eth.M0Listener(ctx, logMessagePublishedMap, log, ws, processingQueue)
			})

			// cleanup irrelevant LogMessagePublished events
			go func() {
				ticker := time.NewTicker(10 * time.Second)
				defer ticker.Stop()
				for range ticker.C {
					logMessagePublishedMap.Cleanup()
				}
			}()

			// TODO: build out wormhole query...
			go func() {
				for {
					dequeued := <-processingQueue
					fmt.Println("READY TO QUERY", dequeued.Sequence)
				}
			}()

			if err := g.Wait(); err != nil {
				return err
			}
			return nil
		},
	}

	appstate.AddConfigurationFlags(cmd)

	return cmd
}
