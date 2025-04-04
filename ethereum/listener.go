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
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"jester.noble.xyz/ethereum/abi/mportal"
	"jester.noble.xyz/ethereum/abi/wormhole"
	"jester.noble.xyz/state"
	"jester.noble.xyz/utils"
)

// subscribeToEvent defines a callback function for subscribing to Ethereum events.
// It establishes a subscription to the desired events, receives a context for cancellation,
// and sends event data to the provided sink channel.
//
// Returns:
// - event.Subscription: The subscription object to manage the subscription lifecycle.
// - error: Any error encountered during the subscription process.
//
// T is the type representing the Ethereum event being subscribed to.
type subscribeToEvent[T any] func(ctx context.Context, sink chan T) (event.Subscription, error)

// processEvent defines a callback function to handle incoming blockchain events from a subscription.
// This function processes a single event.
//
// Note:
// - This function does not return an error to ensure the listener remains operational regardless of processing issues.
// - Any errors encountered should be logged within the implementation.
// - Panics during event processing are recovered within the startEventListener function to ensure resilience.
//
// T is the type representing the Ethereum event being processed.
type processEvent[T any] func(ctx context.Context, log *slog.Logger, event T)

// startEventListener manages the lifecycle of Ethereum event subscriptions and their processing logic.
// This function ensures consistent error handling and resiliency across multiple subscriptions using the
// same websocket client.
//
// Purpose:
// - Establishes a subscription to listen for Ethereum events for a specific contract and event type.
// - Handles errors from the subscription, including re-subscribing or reconnecting the client as needed.
// - Processes events received from the subscription using the provided processEvent callback.
//
// Parameters:
// - ctx (context.Context): Controls the lifecycle of the event listener.
// - log (*slog.Logger): Logger for structured logging, augmented with contract and event details.
// - e (*Eth): The Ethereum client wrapper used for redialing or other utilities.
// - contractName (string): Name of the smart contract (used for logging only).
// - eventName (string): Name of the event being subscribed to (used for logging only).
// - subscribeToEvent (subscribeToEvent[T]): Callback function to establish the event subscription.
// - processEvent (processEvent[T]): Callback function to process incoming events.
//
// Returns:
// - error: An error is only returned if the subscription fails or becomes disconnected and is unrecoverable.
// Returning an error will shutdown Jester.
//
// Example Usage:
//
//	err := startEventListener(
//	    ctx, log, e, "MyContract", "MyEvent",
//	    func(ctx context.Context, sink chan *MyEvent) (event.Subscription, error) {
//			binding, err := myABIContract.NewBindings(common.HexToAddress(myContractAddress), websocketClient)
//			if err != nil {
//				return nil, fmt.Errorf("failed to bind client to myContract: %w", err)
//			}
//	        return binding.WatchMyEvent(opts, sink, filterOptions)
//	    },
//	    func(ctx context.Context, log *slog.Logger, event *MyEvent) {
//	        // Process the event here
//	    },
//	)
//	if err != nil {
//	    log.Error("failed to start event listener", "error", err)
//	}
func startEventListener[T any](
	ctx context.Context, log *slog.Logger,
	e *Eth,
	contractName, eventName string, // used for logging only
	subscribeToEvent subscribeToEvent[T],
	processEvent processEvent[T],
) error {
	log = log.With(slog.String("contract", contractName), slog.String("event", eventName))

	// Channel to receive events/logs from the subscription
	sink := make(chan T)

	// Reconnect loop ensures resiliency in case of subscription errors.
	// Attempts to resubscribe first, if that fails, attempts to redial the Ethereum client.
	for {
		sub, err := subscribeToEvent(ctx, sink)
		if err != nil {
			log.Error("failed to subscribe to events. Attempting to redial client", "error", err)
			err = e.handleRedial(ctx, log)
			if err != nil {
				return errors.Join(errors.New("failed to redial Ethereum client"), err)
			}
			continue
		}
		defer sub.Unsubscribe()

		log.Info("successfully subscribed to events")

		// Process incoming events
	processLoop:
		for {
			select {
			case <-ctx.Done():
				return nil
			case event := <-sink:
				// ensure listener remains running even during panics
				func() {
					defer func() {
						if r := recover(); r != nil {
							log.Error("panic while processing event", "error", r)
						}
					}()
					processEvent(ctx, log, event)
				}()
			case err := <-sub.Err():
				log.Error("subscription error. Attempting to re-subscribe.", "error", err)
				sub.Unsubscribe()
				break processLoop
			}
		}
	}
}

//

// StartWormholeLMPListener listens for Wormhole's `LogMessagePublished` events
func (e *Eth) StartWormholeLMPListener(ctx context.Context, log *slog.Logger, logMessagePublishedMap *state.LogMessagePublishedMap) error {
	contractName := "wormhole"
	eventName := "LogMessagePublished"
	return startEventListener(
		ctx, log, e,
		contractName, eventName,
		func(ctx context.Context, sink chan *wormhole.AbiLogMessagePublished) (event.Subscription, error) {
			wormholeBinding, err := wormhole.NewAbiFilterer(common.HexToAddress(e.Config.WormholeCore), e.WebsocketClient)
			if err != nil {
				return nil, fmt.Errorf("failed to bind client to wormhole contract: %w", err)
			}

			return wormholeBinding.WatchLogMessagePublished(
				&bind.WatchOpts{Context: ctx},
				sink,
				[]common.Address{common.HexToAddress(e.Config.WormholeTransceiver)},
			)
		},
		func(ctx context.Context, log *slog.Logger, event *wormhole.AbiLogMessagePublished) {
			log.Debug("observed event", "txHash", event.Raw.TxHash.String(), "sequence", event.Sequence)
			logMessagePublishedMap.Store(event.Raw.TxHash.String(), event.Sequence)
			e.Metrics.IncLogMessagePublishedCounter()
		},
	)
}

// StartM0TokenSentListener listens for M0's `MTokenSent` events
func (e *Eth) StartM0TokenSentListener(ctx context.Context, log *slog.Logger, logMessagePublishedMap *state.LogMessagePublishedMap, processingQueue chan *utils.QueryData) error {
	contractName := "m0"
	eventName := "MTokenSent"
	return startEventListener(
		ctx, log, e, contractName, eventName,
		func(ctx context.Context, sink chan *mportal.AbiMTokenSent) (event.Subscription, error) {
			binding, err := mportal.NewAbi(common.HexToAddress(e.Config.HubPortal), e.WebsocketClient)
			if err != nil {
				return nil, fmt.Errorf("failed to bind client to mportal contract: %w", err)
			}

			return binding.WatchMTokenSent(
				&bind.WatchOpts{Context: ctx},
				sink,
				nil, nil, nil,
			)
		},
		func(ctx context.Context, log *slog.Logger, event *mportal.AbiMTokenSent) {
			if event.DestinationChainId == e.Config.WormholeNobleChainID {
				processM0Event(ctx, log, event.Raw.TxHash.String(), logMessagePublishedMap, processingQueue)
				e.Metrics.IncMTokenSentCounter()
			}
		},
	)
}

// StartM0MTokenIndexSentListener listens for M0's `MTokenIndexSent` events
func (e *Eth) StartM0MTokenIndexSentListener(ctx context.Context, log *slog.Logger, logMessagePublishedMap *state.LogMessagePublishedMap, processingQueue chan *utils.QueryData) error {
	contractName := "m0"
	eventName := "MTokenIndexSent"
	return startEventListener(
		ctx, log, e, contractName, eventName,
		func(ctx context.Context, sink chan *mportal.AbiMTokenIndexSent) (event.Subscription, error) {
			binding, err := mportal.NewAbi(common.HexToAddress(e.Config.HubPortal), e.WebsocketClient)
			if err != nil {
				return nil, fmt.Errorf("failed to bind client to mportal contract: %w", err)
			}

			return binding.WatchMTokenIndexSent(
				&bind.WatchOpts{Context: ctx},
				sink,
				[]uint16{e.Config.WormholeNobleChainID},
			)
		},
		func(ctx context.Context, log *slog.Logger, event *mportal.AbiMTokenIndexSent) {
			processM0Event(ctx, log, event.Raw.TxHash.String(), logMessagePublishedMap, processingQueue)
			e.Metrics.IncMTokenIndexSentCounter()
		},
	)
}

// processM0Event processes M0 events by looking up the logMessagePublished event that occur in the same transaction.
// All relevant M0 events should have a corresponding logMessagePublished event.
func processM0Event(
	ctx context.Context, log *slog.Logger,
	txHash string,
	logMessagePublishedMap *state.LogMessagePublishedMap,
	processingQueue chan *utils.QueryData,
) {
	log.Info("observed event", "txHash", txHash)

	// Occasionally, the M0 event gets added to state before the LotMessagedPublished event (microseconds difference).
	// To combat this, we retry.
	var seq uint64
	var found bool
	retryAttempts := 5
	err := retry.Do(func() error {
		seq, found = logMessagePublishedMap.GetSequence(txHash)
		if !found {
			return errors.New("not found")
		}
		return nil
	},
		retry.Context(ctx),
		retry.Attempts(uint(retryAttempts)),
		retry.Delay(5*time.Millisecond),
		retry.OnRetry(func(attempt uint, _ error) {
			log.Debug("retry: logMessagePublished lookup", "attempt", fmt.Sprintf("%d/%d", attempt+1, retryAttempts), "txHash", txHash)
		}),
	)
	if err != nil {
		log.Error("`logMessagePublished` sequence not found in correlation to `M0` event", "txHash", txHash)
	}
	if err == nil {
		log.Info("found correlating `logMessagePublished` event", "txHash", txHash, "sequence", seq)

		processingQueue <- &utils.QueryData{
			Sequence: seq,
			TxHash:   txHash,
		}
		logMessagePublishedMap.Delete(txHash)
	}
}
