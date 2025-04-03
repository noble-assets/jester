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
	"fmt"
	"log/slog"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	eth "jester.noble.xyz/ethereum"
	"jester.noble.xyz/wormhole/abi/mportal"
	wormholeabi "jester.noble.xyz/wormhole/abi/wormhole"
)

// StartWormholeLMPListener listens for Wormhole's `LogMessagePublished` events
func (w *Wormhole) StartWormholeLMPListener(ctx context.Context, log *slog.Logger, e *eth.Eth) error {
	contractName := "wormhole"
	eventName := "LogMessagePublished"
	return eth.StartEventListener(
		ctx, log, e,
		contractName, eventName,
		func(ctx context.Context, sink chan *wormholeabi.AbiLogMessagePublished) (event.Subscription, error) {
			wormholeBinding, err := wormholeabi.NewAbiFilterer(common.HexToAddress(w.Config.WormholeCore), e.WebsocketClient)
			if err != nil {
				return nil, fmt.Errorf("failed to bind client to wormhole contract: %w", err)
			}

			return wormholeBinding.WatchLogMessagePublished(
				&bind.WatchOpts{Context: ctx},
				sink,
				[]common.Address{common.HexToAddress(w.Config.WormholeTransceiver)},
			)
		},
		func(ctx context.Context, log *slog.Logger, event *wormholeabi.AbiLogMessagePublished) {
			log.Debug("observed event", "txHash", event.Raw.TxHash.String(), "sequence", event.Sequence)
			w.logMessagePublishedMap.Store(event.Raw.TxHash.String(), event.Sequence)
			w.metrics.IncLogMessagePublishedCounter()
		},
	)
}

// StartM0TokenSentListener listens for M0's `MTokenSent` events
func (w *Wormhole) StartM0TokenSentListener(ctx context.Context, log *slog.Logger, e *eth.Eth) error {
	contractName := "m0"
	eventName := "MTokenSent"
	return eth.StartEventListener(
		ctx, log, e, contractName, eventName,
		func(ctx context.Context, sink chan *mportal.AbiMTokenSent) (event.Subscription, error) {
			binding, err := mportal.NewAbi(common.HexToAddress(w.Config.HubPortal), e.WebsocketClient)
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
			if event.DestinationChainId == w.Config.WormholeNobleChainID {
				w.processM0Event(ctx, log, event.Raw.TxHash.String())
				w.metrics.IncMTokenSentCounter()
			}
		},
	)
}

// StartM0MTokenIndexSentListener listens for M0's `MTokenIndexSent` events
func (w *Wormhole) StartM0MTokenIndexSentListener(ctx context.Context, log *slog.Logger, e *eth.Eth) error {
	contractName := "m0"
	eventName := "MTokenIndexSent"
	return eth.StartEventListener(
		ctx, log, e, contractName, eventName,
		func(ctx context.Context, sink chan *mportal.AbiMTokenIndexSent) (event.Subscription, error) {
			binding, err := mportal.NewAbi(common.HexToAddress(w.Config.HubPortal), e.WebsocketClient)
			if err != nil {
				return nil, fmt.Errorf("failed to bind client to mportal contract: %w", err)
			}

			return binding.WatchMTokenIndexSent(
				&bind.WatchOpts{Context: ctx},
				sink,
				[]uint16{w.Config.WormholeNobleChainID},
			)
		},
		func(ctx context.Context, log *slog.Logger, event *mportal.AbiMTokenIndexSent) {
			w.processM0Event(ctx, log, event.Raw.TxHash.String())
			w.metrics.IncMTokenIndexSentCounter()
		},
	)
}

// processM0Event processes M0 events by looking up the logMessagePublished event that occur in the same transaction.
// All relevant M0 events should have a corresponding logMessagePublished event.
func (w *Wormhole) processM0Event(
	ctx context.Context, log *slog.Logger,
	txHash string,
) {
	log.Info("observed event", "txHash", txHash)

	// Occasionally, the M0 event gets added to state before the LotMessagedPublished event (microseconds difference).
	// To combat this, we retry.
	var seq uint64
	var found bool
	retryAttempts := 5
	err := retry.Do(func() error {
		seq, found = w.logMessagePublishedMap.GetSequence(txHash)
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

		w.processingQueue <- &QueryData{
			Sequence: seq,
			TxHash:   txHash,
		}
		w.logMessagePublishedMap.Delete(txHash)
	}
}
