package ethereum

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"jester.noble.xyz/ethereum/abi/mportal"
	"jester.noble.xyz/ethereum/abi/wormhole"
	"jester.noble.xyz/state"
	"jester.noble.xyz/utils"
)

// StartWormholeListener listens for Wormhole's `LogMessagePublished` events
func (e *Eth) StartWormholeListener(
	ctx context.Context, log *slog.Logger,
	logMessagePublishedMap *state.LogMessagePublishedMap,
) error {
	log = log.With(slog.String("listener", "wormhole"), slog.String("event", "LogMessagePublished"))

	sink := make(chan *wormhole.AbiLogMessagePublished)

	// Subscribe to events via an ABI filter.
	// The for loop ensures we continue to re-subscribe and/or redial the
	// websocket client in case of a subscription error.
	for {
		wormholeBinding, err := wormhole.NewAbiFilterer(common.HexToAddress(e.Config.WormholeCore), e.WebsocketClient)
		if err != nil {
			return fmt.Errorf("failed to bind client to wormhole contract: %w", err)
		}

		sub, err := wormholeBinding.WatchLogMessagePublished(
			&bind.WatchOpts{Context: ctx},
			sink,
			[]common.Address{common.HexToAddress(e.Config.WormholeTransceiver)},
		)
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

		// listen and store data from events
	listenloop:
		for {
			select {
			case <-ctx.Done():
				return nil
			case event := <-sink:
				log.Debug("observed event", "txHash", event.Raw.TxHash.String(), "sequence", event.Sequence)
				logMessagePublishedMap.Store(event.Raw.TxHash.String(), event.Sequence)
			case err := <-sub.Err():
				log.Error("subscription error. Attempting to re-subscribe.", "error", err)
				sub.Unsubscribe()
				break listenloop
			}
		}
	}
}

// StartM0Listener listens for M0's MTokenSent events
func (e *Eth) StartMTokenSentListener(
	ctx context.Context, log *slog.Logger,
	logMessagePublishedMap *state.LogMessagePublishedMap,
	processingQueue chan *utils.QueryData,
) error {
	log = log.With(slog.String("listener", "m0"), slog.String("event", "MTokenSent"))

	sink := make(chan *mportal.BindingsMTokenSent)

	// Subscribe to events via an ABI filter.
	// The for loop ensures we continue to re-subscribe and/or redial the
	// websocket client in case of a subscription error.
	for {
		binding, err := mportal.NewBindings(common.HexToAddress(e.Config.HubPortal), e.WebsocketClient)
		if err != nil {
			return fmt.Errorf("failed to bind client to mportal contract: %w", err)
		}

		var sub event.Subscription
		sub, err = binding.WatchMTokenSent(
			&bind.WatchOpts{Context: ctx},
			sink,
			[]uint16{e.Config.WormholeNobleChainID},
			nil, nil,
		)
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

		// listen and process data from events
	listenloop:
		for {
			select {
			case <-ctx.Done():
				return nil
			case event := <-sink:
				txHash := event.Raw.TxHash.String()
				log.Info("observed `MTokenSent` event", "txHash", txHash)

				processM0Event(ctx, log, txHash, logMessagePublishedMap, processingQueue)

			case err := <-sub.Err():
				log.Error("subscription error. Attempting to re-subscribe.", "error", err)
				sub.Unsubscribe()
				break listenloop
			}
		}
	}
}

// StartM0Listener listens for M0's MTokenSent events
func (e *Eth) StartMTokenIndexSentListener(
	ctx context.Context, log *slog.Logger,
	logMessagePublishedMap *state.LogMessagePublishedMap,
	processingQueue chan *utils.QueryData,
) error {
	log = log.With(slog.String("listener", "m0"), slog.String("event", "MTokenIndexSent"))

	sink := make(chan *mportal.BindingsMTokenIndexSent)

	// Subscribe to events via an ABI filter.
	// The for loop ensures we continue to re-subscribe and/or redial the
	// websocket client in case of a subscription error.
	for {
		binding, err := mportal.NewBindings(common.HexToAddress(e.Config.HubPortal), e.WebsocketClient)
		if err != nil {
			return fmt.Errorf("failed to bind client to mportal contract: %w", err)
		}

		var sub event.Subscription
		sub, err = binding.WatchMTokenIndexSent(
			&bind.WatchOpts{Context: ctx},
			sink,
			[]uint16{e.Config.WormholeNobleChainID},
		)
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

		// listen and process data from events
	listenloop:
		for {
			select {
			case <-ctx.Done():
				return nil
			case event := <-sink:
				txHash := event.Raw.TxHash.String()
				log.Info("observed `MTokenSent` event", "txHash", txHash)

				processM0Event(ctx, log, txHash, logMessagePublishedMap, processingQueue)

			case err := <-sub.Err():
				log.Error("subscription error. Attempting to re-subscribe.", "error", err)
				sub.Unsubscribe()
				break listenloop
			}
		}
	}
}

// processM0Event processes M0 events by looking up the relevant logMessagePublished event tied to it.
// All relevant M0 events should have a corresponding logMessagePublished event.
func processM0Event(
	ctx context.Context, log *slog.Logger,
	txHash string,
	logMessagePublishedMap *state.LogMessagePublishedMap,
	processingQueue chan *utils.QueryData,
) {
	// LogMessagedPublished and (MToken or MTokenIndexSent) sent events happen in the same transaction. We use
	// separate websockets to subscribe to each event. In cases where the MTokenSent event
	// gets added to state before the LotMessagedPublished event, we will be unable to query for
	// the sequence number emitted by LotMessagedPublished. For that reason, we retry.
	var seq uint64
	var found bool
	retryAttemps := 5
	err := retry.Do(func() error {
		seq, found = logMessagePublishedMap.GetSequence(txHash)
		if !found {
			return errors.New("not found")
		}
		return nil
	},
		retry.Context(ctx),
		retry.Attempts(uint(retryAttemps)),
		retry.Delay(5*time.Millisecond),
		retry.OnRetry(func(attempt uint, _ error) {
			log.Debug("retry: logMessagePublished lookup", "attempt", fmt.Sprintf("%d/%d", attempt+1, retryAttemps), "txHash", txHash)
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
	mPortalAbi, err := abi.JSON(strings.NewReader(mportal.BindingsMetaData.ABI))
	if err != nil {
		log.Error("unable to parse MTokenSent ABI when querying history", "error", err)
		return
	}

	nobleChainIDHash := common.BigToHash(big.NewInt(int64(e.Config.WormholeNobleChainID)))

	mTokenSentFuncSig := mPortalAbi.Events["MTokenSent"].ID
	mTokenSentFunclogs, err := e.filterLogs(
		ctx, start, end,
		e.Config.HubPortal,
		[][]common.Hash{{mTokenSentFuncSig}, {nobleChainIDHash}},
	)
	if err != nil {
		log.Error("unable to filter `mTokenSent` logs when querying history", "error", err)
		return
	}

	mTokenIndexSentSig := mPortalAbi.Events["MTokenIndexSent"].ID
	mTokenIndexSentLogs, err := e.filterLogs(
		ctx, start, end,
		e.Config.HubPortal,
		[][]common.Hash{{mTokenIndexSentSig}, {nobleChainIDHash}},
	)
	if err != nil {
		log.Error("unable to filter `mTokenIndexSent` logs when querying history", "error", err)
		return
	}

	allM0Logs := append(mTokenSentFunclogs, mTokenIndexSentLogs...)

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
				log.Debug("found relevant events during historical query", "block", event.Raw.BlockNumber, "seq", event.Sequence)
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
