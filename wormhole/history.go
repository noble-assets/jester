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
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	eth "jester.noble.xyz/ethereum"
	"jester.noble.xyz/wormhole/abi/mportal"
	wormholeabi "jester.noble.xyz/wormhole/abi/wormhole"
)

// GetHistory queries historical data.
//
// Since getting historical data is not crucial to Jester, we do not return an error.
// Instead, we log the error and continue.
func (w *Wormhole) GetHistory(
	ctx context.Context, log *slog.Logger,
	e *eth.Eth,
	startBlock, endBlock int64,
) {
	start := big.NewInt(startBlock)
	var end *big.Int
	if endBlock != 0 {
		end = big.NewInt(endBlock)
	}

	log = log.With(slog.String("interop-framework", "wormhole"), slog.Int64("start-block", startBlock), slog.Int64("end-block", endBlock))

	log.Info("starting historical query", "start-block", startBlock, "end-block", endBlock)

	var totalVaas int
	defer func() {
		log.Info("finished historical query", "vaas-found", totalVaas)
	}()

	// load mPortal ABI to get function signatures
	mPortalAbi, err := abi.JSON(strings.NewReader(mportal.AbiMetaData.ABI))
	if err != nil {
		log.Error("unable to parse MTokenSent ABI when querying history", "error", err)
		return
	}

	mTokenSentLogs, err := e.FilterLogs(
		ctx, log,
		ethereum.FilterQuery{
			FromBlock: start,
			ToBlock:   end,
			Addresses: []common.Address{common.HexToAddress(w.config.hubPortal)},
			Topics:    [][]common.Hash{{mPortalAbi.Events["MTokenSent"].ID}},
		},
	)
	if err != nil {
		log.Error("unable to filter `mTokenSent` logs when querying history", "error", err)
		return
	}

	w.metrics.AddMTokenSentCounter(len(mTokenSentLogs))

	var filteredMTokenSentLogs []ethTypes.Log
	for _, mTokenSentLog := range mTokenSentLogs {
		var event mportal.AbiMTokenSent
		if err := mPortalAbi.UnpackIntoInterface(&event, "MTokenSent", mTokenSentLog.Data); err != nil {
			log.Error("error unpacking portal abi into interface when querying history", "error", err)
		}

		if event.DestinationChainId == w.config.wormholeNobleChainID {
			filteredMTokenSentLogs = append(filteredMTokenSentLogs, mTokenSentLog)
		}
	}

	mTokenIndexSentSig := mPortalAbi.Events["MTokenIndexSent"].ID
	nobleChainIDHash := common.BigToHash(big.NewInt(int64(w.config.wormholeNobleChainID)))
	mTokenIndexSentLogs, err := e.FilterLogs(
		ctx, log,
		ethereum.FilterQuery{
			FromBlock: start,
			ToBlock:   end,
			Addresses: []common.Address{common.HexToAddress(w.config.hubPortal)},
			Topics:    [][]common.Hash{{mTokenIndexSentSig}, {nobleChainIDHash}},
		},
	)
	if err != nil {
		log.Error("unable to filter `mTokenIndexSent` logs when querying history", "error", err)
		return
	}

	w.metrics.AddMTokenIndexSentCounter(len(mTokenIndexSentLogs))

	allM0Logs := append(filteredMTokenSentLogs, mTokenIndexSentLogs...)

	if len(allM0Logs) == 0 {
		return
	}

	// load wormhole ABI and get function signature
	wormholeAbi, err := abi.JSON(strings.NewReader(wormholeabi.AbiABI))
	if err != nil {
		log.Error("unable to parse Wormhole ABI when querying history", "error", err)
		return
	}

	logMessagePublishedFuncSig := wormholeAbi.Events["LogMessagePublished"].ID
	logMessagePublishedLogs, err := e.FilterLogs(
		ctx, log,
		ethereum.FilterQuery{
			FromBlock: start,
			ToBlock:   end,
			Addresses: []common.Address{common.HexToAddress(w.config.wormholeCore)},
			Topics:    [][]common.Hash{{logMessagePublishedFuncSig}, {common.HexToHash(w.config.wormholeTransceiver)}},
		},
	)
	if err != nil {
		log.Error("unable to filter `logMessagePublished` logs when querying history", "error", err)
		return
	}

	w.metrics.AddLogMessagePublishedCounter(len(logMessagePublishedLogs))

	event := struct {
		wormholeabi.AbiLogMessagePublished
	}{}

	for _, l := range allM0Logs {
		txHash := l.TxHash
		for _, lLog := range logMessagePublishedLogs {
			if txHash == lLog.TxHash {
				if err := wormholeAbi.UnpackIntoInterface(&event, "LogMessagePublished", lLog.Data); err != nil {
					log.Error("error unpacking wormhole abi into interface when querying history", "error", err)
				}
				log.Debug("found relevant events during historical query", "block", lLog.BlockNumber, "seq", event.Sequence)
				w.processingQueue <- &queryData{
					sequence: event.Sequence,
					txHash:   txHash.String(),
				}
				totalVaas += 1
			}
		}
	}
}
