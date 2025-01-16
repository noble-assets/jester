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
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"jester.noble.xyz/ethereum/abi/mportal"
	"jester.noble.xyz/ethereum/abi/wormhole"
	"jester.noble.xyz/state"
	"jester.noble.xyz/utils"
)

// StartWormholeListener listens for `LogMessagePublished` events
func (e *Eth) StartWormholeListener(
	ctx context.Context, log *slog.Logger,
	logMessagePublishedMap *state.LogMessagePublishedMap,
) error {
	log = log.With(slog.String("listener", "wormhole"))

	sink := make(chan *wormhole.AbiLogMessagePublished)

	// Subscribe to LogMessagePublished events.
	// The for loop ensures we continue to re-subscribe and/or redial the
	// websocket client in case of a subscription error.
	for {
		backend := e.NewWebsocketContractBackendWrapper()
		wormholeBinding, err := wormhole.NewAbiFilterer(common.HexToAddress(e.Config.WormholeCore), backend)
		if err != nil {
			return fmt.Errorf("failed to bind client to wormhole contract: %w", err)
		}

		sub, err := wormholeBinding.WatchLogMessagePublished(
			&bind.WatchOpts{Context: ctx},
			sink,
			[]common.Address{common.HexToAddress(e.Config.WormholeTransceiver)},
		)
		if err != nil {
			log.Error("failed to subscribe to `LogMessagePublished` events. Attempting to redial client", "error", err)
			err = e.handleRedial(ctx, log)
			if err != nil {
				return errors.Join(errors.New("failed to redial Ethereum client"), err)
			}
			continue
		}
		defer sub.Unsubscribe()

		log.Info("successfully subscribed to `LogMessagePublished` events")

		// listen and store data from events
	listenloop:
		for {
			select {
			case <-ctx.Done():
				return nil
			case event := <-sink:
				log.Debug("observed `LogMessagePublished` event", "txHash", event.Raw.TxHash.String(), "sequence", event.Sequence)
				logMessagePublishedMap.Store(event.Raw.TxHash.String(), event.Sequence)
			case err := <-sub.Err():
				log.Error("subscription error. Attempting to re-subscribe.", "error", err)
				sub.Unsubscribe()
				break listenloop
			}
		}
	}
}

// StartM0Listener listens for MTokenSent events
func (e *Eth) StartM0Listener(
	ctx context.Context, log *slog.Logger,
	logMessagePublishedMap *state.LogMessagePublishedMap,
	processingQueue chan *utils.QueryData,
) error {
	log = log.With(slog.String("listener", "m0"))

	sink := make(chan *mportal.BindingsMTokenSent)

	// Subscribe to MTokenSent events.
	// The for loop ensures we continue to re-subscribe and/or redial the
	// websocket client in case of a subscription error.
	for {
		backend := e.NewWebsocketContractBackendWrapper()
		binding, err := mportal.NewBindings(common.HexToAddress(e.Config.HubPortal), backend)
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
			log.Error("failed to subscribe to `MTokenSent` events. Attempting to redial client", "error", err)
			err = e.handleRedial(ctx, log)
			if err != nil {
				return errors.Join(errors.New("failed to redial Ethereum client"), err)
			}
			continue
		}
		defer sub.Unsubscribe()

		log.Info("successfully subscribed to `MTokenSent` events")

		// listen and process data from events
	listenloop:
		for {
			select {
			case <-ctx.Done():
				return nil
			case event := <-sink:
				txHash := event.Raw.TxHash.String()
				log.Info("observed `MTokenSent` event", "txHash", txHash)

				// LogMessagedPublished and MToken sent events happen in the same transaction. We use
				// separate websockets to subscribe to each event. In cases where we happen observe the
				// MTokenSent event before the LotMessagedPublished event, we will be unable to query for
				// the sequence number emitted by LotMessagedPublished. For that reason, we retry.
				var seq uint64
				var found bool
				retryAttemps := 3
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
					log.Error("`logMessagePublished` sequence not found in correlation to `MTokenSent` event", "txHash", txHash)
				}
				if err == nil {
					log.Info("found correlating `logMessagePublished` event", "txHash", txHash, "sequence", seq)

					processingQueue <- &utils.QueryData{
						WormHoleChainID: e.Config.WormholeSrcChainId,
						Emitter:         e.Config.WormholeTransceiver,
						Sequence:        seq,
						TxHash:          txHash,
					}
					logMessagePublishedMap.Delete(event.Raw.TxHash.String())
				}
			case err := <-sub.Err():
				log.Error("subscription error. Attempting to re-subscribe.", "error", err)
				sub.Unsubscribe()
				break listenloop
			}
		}
	}
}

// GetHistory queries historical data.
func (e *Eth) GetHistory(
	ctx context.Context, log *slog.Logger,
	processingQueue chan *utils.QueryData,
	startBlock int64, endBlock int64,
) {
	from := big.NewInt(startBlock)
	var end *big.Int
	if endBlock != 0 {
		end = big.NewInt(endBlock)
	}

	log = log.With(slog.Int64("start-block", startBlock), slog.Int64("end-block", endBlock))
	log.Info("starting to query history")

	rpc := e.RPCClient

	var totalVaas int
	defer func() {
		log.Info("finished querying history", "vaas-found", totalVaas)
	}()

	// load mPortal ABI and get MTokenSent function signature
	mPortalAbi, err := abi.JSON(strings.NewReader(mportal.BindingsMetaData.ABI))
	if err != nil {
		log.Error("unable to parse MTokenSent ABI when querying history", "error", err)
		return
	}
	mTokenSentFuncSig := mPortalAbi.Events["MTokenSent"].ID

	nobleChainIDHash := common.BigToHash(big.NewInt(int64(e.Config.WormholeNobleChainID)))

	query := ethereum.FilterQuery{
		FromBlock: from,
		ToBlock:   end,
		Addresses: []common.Address{common.HexToAddress(e.Config.HubPortal)},
		Topics:    [][]common.Hash{{mTokenSentFuncSig}, {nobleChainIDHash}},
	}

	mTokenSentLogs, err := rpc.FilterLogs(ctx, query)
	if err != nil {
		log.Error("unable to filter `mTokenSent` logs when querying history", "error", err)
		return
	}

	if len(mTokenSentLogs) == 0 {
		return
	}

	// load wormhole ABI and get LogMessagePublished function signature
	wormholeAbi, err := abi.JSON(strings.NewReader(wormhole.AbiABI))
	if err != nil {
		log.Error("unable to parse Wormhole ABI when querying history", "error", err)
		return
	}
	logMessagePublishedFuncSig := wormholeAbi.Events["LogMessagePublished"].ID

	// update query
	query.Addresses = []common.Address{common.HexToAddress(e.Config.WormholeCore)}
	query.Topics = [][]common.Hash{{logMessagePublishedFuncSig}, {common.HexToHash(e.Config.WormholeTransceiver)}}

	logMessagePublishedLogs, err := rpc.FilterLogs(ctx, query)
	if err != nil {
		log.Error("unable to filter `logMessagePublished` logs when querying history", "error", err)
		return
	}

	event := struct {
		wormhole.AbiLogMessagePublished
	}{}

	for _, mLog := range mTokenSentLogs {
		txHash := mLog.TxHash
		for _, lLog := range logMessagePublishedLogs {
			if txHash == lLog.TxHash {
				if err := wormholeAbi.UnpackIntoInterface(&event, "LogMessagePublished", lLog.Data); err != nil {
					log.Error("error unpacking wormhole abi into interface when querying history", "error", err)
				}
				log.Debug("found relevant events during historical query", "block", event.Raw.BlockNumber, "seq", event.Sequence)
				processingQueue <- &utils.QueryData{
					WormHoleChainID: e.Config.WormholeSrcChainId,
					Emitter:         e.Config.WormholeTransceiver,
					Sequence:        event.Sequence,
					TxHash:          txHash.String(),
				}
				totalVaas += 1
			}
		}
	}
}

type ContractBackendWrapper struct {
	*ethclient.Client
}

// NewWebsocketContractBackendWrapper creates the backend for a websocket client.
// This is used to bind to an abi filter.
func (e *Eth) NewWebsocketContractBackendWrapper() *ContractBackendWrapper {
	e.WebsocketClientMutex.Lock()
	defer e.WebsocketClientMutex.Unlock()

	return &ContractBackendWrapper{Client: e.WebsocketClient}
}
