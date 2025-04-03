package wormhole

import (
	"context"
	"log/slog"
	"math/big"
	"strings"

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
	startBlock int64, endBlock int64,
) {
	start := big.NewInt(startBlock)
	var end *big.Int
	if endBlock != 0 {
		end = big.NewInt(endBlock)
	}

	log = log.With(slog.Int64("start-block", startBlock), slog.Int64("end-block", endBlock))
	log.Info("starting to query wormhole related history")

	var totalVaas int
	defer func() {
		log.Info("finished querying wormhole related history", "vaas-found", totalVaas)
	}()

	// load mPortal ABI to get function signatures
	mPortalAbi, err := abi.JSON(strings.NewReader(mportal.AbiMetaData.ABI))
	if err != nil {
		log.Error("unable to parse MTokenSent ABI when querying history", "error", err)
		return
	}

	mTokenSentLogs, err := e.FilterLogs(
		ctx, start, end,
		w.Config.HubPortal,
		[][]common.Hash{{mPortalAbi.Events["MTokenSent"].ID}},
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

		if event.DestinationChainId == w.Config.WormholeNobleChainID {
			filteredMTokenSentLogs = append(filteredMTokenSentLogs, mTokenSentLog)
		}
	}

	mTokenIndexSentSig := mPortalAbi.Events["MTokenIndexSent"].ID
	nobleChainIDHash := common.BigToHash(big.NewInt(int64(w.Config.WormholeNobleChainID)))
	mTokenIndexSentLogs, err := e.FilterLogs(
		ctx, start, end,
		w.Config.HubPortal,
		[][]common.Hash{{mTokenIndexSentSig}, {nobleChainIDHash}},
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
		ctx, start, end,
		w.Config.WormholeCore,
		[][]common.Hash{{logMessagePublishedFuncSig}, {common.HexToHash(w.Config.WormholeTransceiver)}},
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
				w.processingQueue <- &QueryData{
					Sequence: event.Sequence,
					TxHash:   txHash.String(),
				}
				totalVaas += 1
			}
		}
	}
}
