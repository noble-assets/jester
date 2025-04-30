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

package ethereum

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/sync/semaphore"
)

const (
	// maxConcurrentFinalityChecks is the maximum number of concurrent finality checks
	maxConcurrentFinalityChecks = 100

	// finalityPollInterval is how often to poll for new finalized blocks
	finalityPollInterval = 30 * time.Second
)

// RouteAfterFinality is a callback function that is called after an Ethereum event has been finalized.
// It should pass the Ethereum log to its respective processing queue.
//
// Example:
//
//	f := func(ethLog ethTypes.Log) {
//	    processingQueue <- &ethLog
//	}
type RouteAfterFinality func(ethTypes.Log)

// ethEventForFinality is used to track Ethereum events that need to be finalized.
type ethEventForFinality struct {
	ethLogs ethTypes.Log

	routeAfterFinality RouteAfterFinality

	timeObserved time.Time
}

// EnsureFinality sends an Ethereum event to the finality process.
// This is used to ensure that the event is finalized before processing it.
// Once finalized the event is passed to the RouteAfterFinality callback.
//
// The finality check includes:
// 1. Verifying the event's block is part of the canonical chain
// 2. Checking if the block number is below the current finalized block height
// 3. Attempting to recover the event if a reorg is detected
func (e *Eth) EnsureFinality(ethLog ethTypes.Log, routeAfterFinality RouteAfterFinality) {
	e.ensureFinalityCh <- &ethEventForFinality{
		ethLogs:            ethLog,
		routeAfterFinality: routeAfterFinality,
		timeObserved:       time.Now(),
	}
}

// startFinalityRoutine watches and handles new ethereum events that need to be finalized.
func (e *Eth) startFinalityRoutine(ctx context.Context, log *slog.Logger) error {
	sem := semaphore.NewWeighted(maxConcurrentFinalityChecks)
	errChan := make(chan error, 1)

	for {
		select {
		case <-ctx.Done():
			return nil
		case dequeued, ok := <-e.ensureFinalityCh:
			if !ok {
				return errors.New("ensureFinalityProcess channel closed unexpectedly")
			}

			if err := sem.Acquire(ctx, 1); err != nil {
				return errors.Join(errors.New("failed to acquire semaphore"), err)
			}

			go func(event *ethEventForFinality) {
				defer sem.Release(1)
				if err := e.waitForFinality(ctx, log, event); err != nil {
					errChan <- err
				}
			}(dequeued)
		case err := <-errChan:
			return fmt.Errorf("finality routine error: %w", err)
		}
	}
}

// waitForFinality checks and/or waits for the event to be finalized.
// Returning an error will ultimately cause Jester to quit.
func (e *Eth) waitForFinality(ctx context.Context, log *slog.Logger, event *ethEventForFinality) error {

	// Helper function to check finality and handle the result
	// Returning true indicates the event is either:
	// 	a) finalized and the event has been passed to the callback for processing.
	// 	b) lost in a re-org and the event should be ignored.
	checkAndHandleFinality := func(recordMetric bool) (bool, error) {
		finalized, lostInReorg, newLog, err := e.checkForFinality(ctx, log, event)
		if err != nil {
			return false, fmt.Errorf("failed to check for finality: %w", err)
		}
		event.ethLogs = newLog

		if lostInReorg {
			return true, nil
		}

		if finalized {
			event.routeAfterFinality(event.ethLogs)
			if recordMetric && event.timeObserved != (time.Time{}) {
				e.metrics.ObserveFinalityReachedSeconds(time.Since(event.timeObserved).Seconds())
			}
			return true, nil
		}

		return false, nil
	}

	// First check if already finalized (don't record metric for historical events)
	done, err := checkAndHandleFinality(false)
	if err != nil || done {
		return err
	}

	// If not finalized, subscribe to finalized blocks and wait
	subFinalizedBlock, unsub := e.subToFinalizedBlocks(ctx, log)
	defer unsub()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-subFinalizedBlock:
			done, err := checkAndHandleFinality(true)
			if err != nil || done {
				return err
			}
		}
	}
}

// checkForFinality checks if an Ethereum event is finalized by querying the latest finalized block.
// It also checks for re-orgs and attempts to recover the event if necessary.
//
// Because of the potential for a re-org, we must ensure we return new the ethLog as this log will be the
// new source of truth.
func (e *Eth) checkForFinality(ctx context.Context, log *slog.Logger, event *ethEventForFinality) (finalized, lostInReorg bool, ethLog ethTypes.Log, err error) {
	// first check for re-org by comparing the block hash of the event with the canonical block hash.
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	canonicalBlock, err := e.headerByNumber(ctxWithTimeout, log, big.NewInt(int64(event.ethLogs.BlockNumber)))
	if err != nil {
		return false, false, ethTypes.Log{}, err
	}

	if canonicalBlock.Hash() != event.ethLogs.BlockHash {
		log.Info("re-org detected!",
			"txHash", event.ethLogs.TxHash.Hex(),
			"observed-height", event.ethLogs.BlockNumber,
			"expected-block-hash", event.ethLogs.BlockHash.Hex(),
			"received-block-hash", canonicalBlock.Hash().Hex(),
			"timer", time.Since(event.timeObserved).Round(time.Second).String(),
		)

		e.metrics.IncReorgDetectedCounter()

		success, newLog, err := e.attemptEventRecovery(ctx, log, event.ethLogs)
		if err != nil {
			return false, false, ethTypes.Log{}, fmt.Errorf("failed to attempt event recovery: %w", err)
		}

		if !success {
			return false, true, ethTypes.Log{}, nil
		}

		event.ethLogs = newLog
	}

	// then check if the event is finalized by comparing the block number with the current finalized block.
	finalizedHeight := e.finalizedHeightPoller.currentFinalizedHeight.Load()

	if finalizedHeight >= event.ethLogs.BlockNumber {
		log.Info("finality reached",
			"txHash", event.ethLogs.TxHash.Hex(),
			"observed-block", event.ethLogs.BlockNumber,
			"current-finalized-block", finalizedHeight,
			"timer", time.Since(event.timeObserved).Round(time.Second).String(),
		)
		return true, false, event.ethLogs, nil
	}

	log.Info("waiting for finality",
		"txHash", event.ethLogs.TxHash.Hex(),
		"observed-height", event.ethLogs.BlockNumber,
		"current-finalized-height", finalizedHeight,
		"timer", time.Since(event.timeObserved).Round(time.Second).String(),
	)

	return false, false, event.ethLogs, nil
}

// attemptEventRecovery attempts to recover an event in the event of an Ethereum re-org.
func (e *Eth) attemptEventRecovery(ctx context.Context, log *slog.Logger, ethLog ethTypes.Log) (success bool, newLogs ethTypes.Log, err error) {
	fromBlock := big.NewInt(int64(ethLog.BlockNumber))

	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	logs, err := e.FilterLogs(
		ctxWithTimeout, log,
		ethereum.FilterQuery{
			FromBlock: fromBlock,
			ToBlock:   nil,
			Addresses: []common.Address{ethLog.Address},
			Topics:    [][]common.Hash{ethLog.Topics},
		},
	)
	if err != nil {
		return false, ethTypes.Log{}, fmt.Errorf("failed to filter logs in event recovery: %w", err)
	}

	for _, vLog := range logs {
		if vLog.TxHash == ethLog.TxHash {
			log.Info("recovered event",
				"txHash", vLog.TxHash.Hex(),
				"observed-height", ethLog.BlockNumber,
				"new-height", vLog.BlockNumber)

			e.metrics.IncReorgEventRecoveredCounter()

			return true, vLog, nil
		}
	}

	log.Info("failed to recover event",
		"txHash", ethLog.TxHash.String(),
		"observed-height", ethLog.BlockNumber,
	)

	e.metrics.IncReorgEventLostCounter()

	return false, ethTypes.Log{}, nil
}

/////////////////////
// Finality Poller //
/////////////////////

// finalizedHeightPoller tracks the finalized block height and notifies subscribers on updates.
type finalizedHeightPoller struct {
	currentFinalizedHeight atomic.Uint64
	mu                     sync.Mutex

	// subscribers holds all active subscription channels.
	subscribers map[chan struct{}]struct{}

	// done signals the polling routine that there are no more subscribers and can stop polling.
	done    chan struct{}
	running bool
	errChan chan error
}

// newFinalizedHeightPoller creates a new poller.
func newFinalizedHeightPoller() *finalizedHeightPoller {
	return &finalizedHeightPoller{
		subscribers: make(map[chan struct{}]struct{}),
		done:        make(chan struct{}),
		errChan:     make(chan error, 1),
	}
}

// subToFinalizedBlocks establishes a subscription to receive updates on the finalized block height.
// It returns a buffered channel that emits a notification each time a new finalized block is detected,
// along with an unsubscribe function to cancel the subscription.
// You can get the new finalized block height by calling e.finalizedHeightPoller.currentFinalizedHeight.Load().
//
// Since getting the finalized block is only available via RPC, using this subscription method and only polling
// when necessary helps reduce the amount of RPC calls needed.
func (e *Eth) subToFinalizedBlocks(ctx context.Context, log *slog.Logger) (newFinalizedHeight <-chan struct{}, unsubscribe func()) {
	p := e.finalizedHeightPoller
	p.mu.Lock()
	defer p.mu.Unlock()

	ch := make(chan struct{}, 1)
	p.subscribers[ch] = struct{}{}

	if !p.running {
		p.running = true
		go e.pollFinalizedHeight(ctx, log)
	}

	unsubscribe = func() {
		p.mu.Lock()
		defer p.mu.Unlock()
		if _, exists := p.subscribers[ch]; exists {
			delete(p.subscribers, ch)
			close(ch)
			if len(p.subscribers) == 0 && p.running {
				close(p.done)
				p.running = false
				p.done = make(chan struct{})
			}
		}
	}

	return ch, unsubscribe
}

// pollFinalizedHeight queries for the finalized height every 30 seconds.
func (e *Eth) pollFinalizedHeight(ctx context.Context, log *slog.Logger) {
	log.Info("starting finalized height poller")
	p := e.finalizedHeightPoller
	ticker := time.NewTicker(finalityPollInterval)
	defer ticker.Stop()

	updateHeight := func() {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, finalityPollInterval/2)
		defer cancel()
		header, err := e.headerByNumber(ctxWithTimeout, log, big.NewInt(rpc.FinalizedBlockNumber.Int64()))
		if err != nil {
			p.errChan <- err
			return
		}

		newFinalized := header.Number.Uint64()
		currentFinalized := p.currentFinalizedHeight.Load()

		if newFinalized > currentFinalized {
			p.currentFinalizedHeight.Store(newFinalized)
			log.Info("finalized height updated", "height", newFinalized)

			p.mu.Lock()
			for sub := range p.subscribers {
				// Non-blocking send to avoid blocking the poller.
				select {
				case sub <- struct{}{}:
				default:
				}
			}
			p.mu.Unlock()
		}
	}

	updateHeight()

	for {
		p.mu.Lock()
		doneCh := p.done
		p.mu.Unlock()

		select {
		case <-ctx.Done():
			return
		case <-doneCh:
			log.Info("stopping finalized height poller; all observed events finalized")
			return
		case <-ticker.C:
			updateHeight()
		}
	}
}

// watchForFinalizedHeightPollerErrors starts a goroutine to watch for errors from the finalize block height poller.
func (e *Eth) watchForFinalizedHeightPollerErrors(ctx context.Context) error {
	p := e.finalizedHeightPoller

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-p.errChan:
			return fmt.Errorf("finalized height poller error: %w", err)
		}
	}
}
