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
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	eth "jester.noble.xyz/v2/ethereum"
	hyperlaneabi "jester.noble.xyz/v2/hyperlane/abi"
)

func (h *Hyperlane) GetHistory(
	ctx context.Context, log *slog.Logger,
	e *eth.Eth,
	startBlock, endBlock int64,
) {
	start := big.NewInt(startBlock)
	var end *big.Int
	if endBlock != 0 {
		end = big.NewInt(endBlock)
	}

	log = log.With(slog.String("interop-framework", "hyperlane"), slog.Int64("start-block", startBlock), slog.Int64("end-block", endBlock))

	log.Info("starting historical query")

	mailboxAbi, err := abi.JSON(strings.NewReader(hyperlaneabi.HyperlaneMetaData.ABI))
	if err != nil {
		log.Error("unable to parse Mailbox ABI when querying history", "error", err)
		return
	}

	dispatchLogs, err := e.FilterLogs(
		ctx, log,
		ethereum.FilterQuery{
			FromBlock: start,
			ToBlock:   end,
			Addresses: []common.Address{common.HexToAddress(h.config.hyperlaneMailbox)},
			Topics:    [][]common.Hash{{mailboxAbi.Events["Dispatch"].ID}}, // TODO: filter for noble
		},
	)
	if err != nil {
		log.Error("unable to filter `dispatch` logs when querying history", "error", err)
		return
	}

	h.metrics.AddMailboxDispatchCounter(len(dispatchLogs))

	for _, log := range dispatchLogs {
		e.EnsureFinality(log, h.hyperlaneDispatchRouter)
	}

	log.Info("finished historical query", "events-found", len(dispatchLogs))
}
