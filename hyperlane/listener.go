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

package hyperlane

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	eth "jester.noble.xyz/ethereum"
	hyperlaneabi "jester.noble.xyz/hyperlane/abi"
)

// StartHyperlaneDispatchListener listens for Hyperlane's `Dispatch` events
func (h *Hyperlane) startHyperlaneDispatchListener(
	ctx context.Context,
	log *slog.Logger,
	e *eth.Eth,
) error {
	contractName := "mailbox"
	eventName := "Dispatch"
	return eth.StartEventListener(
		ctx, log, e,
		contractName, eventName,
		func(ctx context.Context, sink chan *hyperlaneabi.HyperlaneDispatch) (event.Subscription, error) {
			binding, err := hyperlaneabi.NewHyperlaneFilterer(common.HexToAddress(h.config.hyperlaneMailbox), e.WebsocketClient)
			if err != nil {
				return nil, fmt.Errorf("failed to bind client to hyperlane mailbox contract: %w", err)
			}

			return binding.WatchDispatch(
				&bind.WatchOpts{Context: ctx},
				sink,
				nil,
				[]uint32{h.config.hyperlaneNobleChainId},
				nil,
			)
		},
		func(ctx context.Context, log *slog.Logger, event *hyperlaneabi.HyperlaneDispatch) {
			log.Info("observed event", "txHash", event.Raw.TxHash.String())
			// TODO: process event
			h.metrics.IncMailboxDispatchCounter()
		},
	)
}
