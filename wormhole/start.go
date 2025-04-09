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

package wormhole

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	eth "jester.noble.xyz/ethereum"
	"jester.noble.xyz/metrics"
)

// Start starts Wormhole bridge integration.
func (w *Wormhole) Start(
	ctx context.Context, log *slog.Logger, Eth *eth.Eth,
	metrics *metrics.PrometheusMetrics,
) error {
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

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return w.StartWormholeLMPListener(ctx, log, Eth)
	})

	g.Go(func() error {
		return w.StartM0TokenSentListener(ctx, log, Eth)
	})

	g.Go(func() error {
		return w.StartM0MTokenIndexSentListener(ctx, log, Eth)
	})

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
				w.logMessagePublishedMap.Cleanup()
			}
		}
	}()

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
			case dequeued, ok := <-w.processingQueue:
				if !ok {
					return errors.New("processingQueue channel closed unexpectedly")
				}
				if err := sem.Acquire(ctx, 1); err != nil {
					return errors.Join(errors.New("failed to acquire semaphore"), err)
				}

				go func(data *queryData) {
					defer sem.Release(1)
					w.StartWormholeWorker(ctx, log, metrics, data)
				}(dequeued)
			}
		}
	})

	// Hold here, waiting for any errors from errgroup
	if err := g.Wait(); err != nil {
		log.Error("wormhole error", "error", err)
		return err
	}

	return nil
}
