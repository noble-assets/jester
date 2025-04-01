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
	"log/slog"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"jester.noble.xyz/ethereum/abi/mportal"
	"jester.noble.xyz/ethereum/abi/wormhole"
	"jester.noble.xyz/utils"
)

// GetHistory queries historical data.
//
// Since getting historical data is not crucial to Jester, we do not return an error.
// Instead, we log the error and continue.
func (e *Eth) GetHistory(
	ctx context.Context, log *slog.Logger,
	processingQueue chan *utils.QueryData,
	startBlock int64, endBlock int64,
) {
	start := big.NewInt(startBlock)
	var end *big.Int
	if endBlock != 0 {
		end = big.NewInt(endBlock)
	}

	log = log.With(slog.Int64("start-block", startBlock), slog.Int64("end-block", endBlock))
	log.Info("starting to query history")

	var totalVaas int
	defer func() {
		log.Info("finished querying history", "vaas-found", totalVaas)
	}()

	// load mPortal ABI to get function signatures
	mPortalAbi, err := abi.JSON(strings.NewReader(mportal.AbiMetaData.ABI))
	if err != nil {
		log.Error("unable to parse MTokenSent ABI when querying history", "error", err)
		return
	}

	mTokenSentLogs, err := e.filterLogs(
		ctx, start, end,
		e.Config.HubPortal,
		[][]common.Hash{{mPortalAbi.Events["MTokenSent"].ID}},
	)
	if err != nil {
		log.Error("unable to filter `mTokenSent` logs when querying history", "error", err)
		return
	}

	e.Metrics.AddMTokenSentCounter(len(mTokenSentLogs))

	var filteredMTokenSentLogs []ethTypes.Log
	for _, mTokenSentLog := range mTokenSentLogs {
		var event mportal.AbiMTokenSent
		if err := mPortalAbi.UnpackIntoInterface(&event, "MTokenSent", mTokenSentLog.Data); err != nil {
			log.Error("error unpacking portal abi into interface when querying history", "error", err)
		}

		if event.DestinationChainId == e.Config.WormholeNobleChainID {
			filteredMTokenSentLogs = append(filteredMTokenSentLogs, mTokenSentLog)
		}
	}

	mTokenIndexSentSig := mPortalAbi.Events["MTokenIndexSent"].ID
	nobleChainIDHash := common.BigToHash(big.NewInt(int64(e.Config.WormholeNobleChainID)))
	mTokenIndexSentLogs, err := e.filterLogs(
		ctx, start, end,
		e.Config.HubPortal,
		[][]common.Hash{{mTokenIndexSentSig}, {nobleChainIDHash}},
	)
	if err != nil {
		log.Error("unable to filter `mTokenIndexSent` logs when querying history", "error", err)
		return
	}

	e.Metrics.AddMTokenIndexSentCounter(len(mTokenIndexSentLogs))

	allM0Logs := append(filteredMTokenSentLogs, mTokenIndexSentLogs...)

	if len(allM0Logs) == 0 {
		return
	}

	// load wormhole ABI and get function signature
	wormholeAbi, err := abi.JSON(strings.NewReader(wormhole.AbiABI))
	if err != nil {
		log.Error("unable to parse Wormhole ABI when querying history", "error", err)
		return
	}

	logMessagePublishedFuncSig := wormholeAbi.Events["LogMessagePublished"].ID
	logMessagePublishedLogs, err := e.filterLogs(
		ctx, start, end,
		e.Config.WormholeCore,
		[][]common.Hash{{logMessagePublishedFuncSig}, {common.HexToHash(e.Config.WormholeTransceiver)}},
	)
	if err != nil {
		log.Error("unable to filter `logMessagePublished` logs when querying history", "error", err)
		return
	}

	e.Metrics.AddLogMessagePublishedCounter(len(logMessagePublishedLogs))

	event := struct {
		wormhole.AbiLogMessagePublished
	}{}

	for _, l := range allM0Logs {
		txHash := l.TxHash
		for _, lLog := range logMessagePublishedLogs {
			if txHash == lLog.TxHash {
				if err := wormholeAbi.UnpackIntoInterface(&event, "LogMessagePublished", lLog.Data); err != nil {
					log.Error("error unpacking wormhole abi into interface when querying history", "error", err)
				}
				log.Debug("found relevant events during historical query", "block", lLog.BlockNumber, "seq", event.Sequence)
				processingQueue <- &utils.QueryData{
					Sequence: event.Sequence,
					TxHash:   txHash.String(),
				}
				totalVaas += 1
			}
		}
	}
}

// filterLogs uses an RPC client to query Ethereum within a specified block range.
// It returns filtered logs based on contract address and topics.
func (e *Eth) filterLogs(ctx context.Context, start, end *big.Int, contractAddress string, topics [][]common.Hash) ([]ethTypes.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: start,
		ToBlock:   end,
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		Topics:    topics,
	}
	return e.RPCClient.FilterLogs(ctx, query)
}
