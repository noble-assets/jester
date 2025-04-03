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
	"log/slog"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/event"
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
func StartEventListener[T any](
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

// handleRedial handles the redial of the websocket client between multiple websocket subscriptions.
// Because the websocket client is shared between multiple subscriptions, this function
// is used to ensure that only one redial is in progress at a time.
func (e *Eth) handleRedial(ctx context.Context, log *slog.Logger) (err error) {
	redial := e.Redial
	redial.inProgressMutex.Lock()

	// Another goroutine is already handling redial
	if redial.inProgress {
		redial.inProgressMutex.Unlock()
		log.Info("client redial already in progress")
		redial.cond.L.Lock()   // Lock mutex to call wait()
		redial.cond.Wait()     // Wait for the redial to complete; unlocks mutex, waits for Broadcast(), re-locks mutex
		redial.cond.L.Unlock() // Unlock mutex
		errExists := false
		if redial.err != nil {
			errExists = true
		}
		log.Info("received client redial complete", "error-exists", errExists)
		return redial.err
	}

	// Mark redial as in progress and prepare a new signal channel
	redial.inProgress = true
	redial.cond = sync.NewCond(&redial.inProgressMutex)
	redial.inProgressMutex.Unlock()

	defer func() {
		redial.inProgressMutex.Lock()
		redial.inProgress = false
		redial.err = err
		redial.cond.Broadcast() // Signal all waiting goroutines that the redial operation is complete
		redial.inProgressMutex.Unlock()
	}()

	e.metrics.IncEthSubInterruptionCounter()

	client, err := dialClient(ctx, log, e.config.websocketURL, websocketClientType)
	if err != nil {
		return err
	}
	e.websocketClientMutex.Lock()
	e.WebsocketClient = client
	e.websocketClientMutex.Unlock()

	time.Sleep(1 * time.Second)     // Allow other redial signals to accumulate
	redial.GetHistory <- struct{}{} // Trigger historical lookup to catch up on missed events
	return nil
}
