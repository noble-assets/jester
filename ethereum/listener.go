package ethereum

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/noble-assets/jester/ethereum/abi/mportal"
	"github.com/noble-assets/jester/ethereum/abi/wormhole"
)

type LogMessagePublishedMap struct {
	mu    sync.Mutex
	store map[string]logEntry // txHash -> logEntry
}

type logEntry struct {
	sequence  uint64
	timestamp time.Time // used to cleanup old entires
}

func NewLogMessagePublishedMap() *LogMessagePublishedMap {
	return &LogMessagePublishedMap{
		store: make(map[string]logEntry),
	}
}

func (m *LogMessagePublishedMap) Store(txHash string, sequence uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.store[txHash] = logEntry{
		sequence:  sequence,
		timestamp: time.Now(),
	}
}

func (m *LogMessagePublishedMap) GetSequence(txHash string) (seq uint64, found bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	entry, found := m.store[txHash]
	if !found {
		return 0, found
	}
	return entry.sequence, found
}

func (m *LogMessagePublishedMap) Delete(txHash string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.store, txHash)
}

func (m *LogMessagePublishedMap) Cleanup() {
	m.mu.Lock()
	defer m.mu.Unlock()
	now := time.Now()
	for key, entry := range m.store {
		if now.Sub(entry.timestamp) > 10*time.Second {
			delete(m.store, key)
		}
	}
}

//

// WormholeListener listens for `LogMessagePublished` events
func WormholeListener(
	ctx context.Context, log *slog.Logger,
	logMessagePublishedMap *LogMessagePublishedMap,
	ws *ethclient.Client,
	wormholeCoreContract, wormholeTransceiverContract string,
) error {
	log = log.With(slog.String("listener", "wormhole"))

	backend := NewContractBackendWrapper(ws)
	wormholeBinding, err := wormhole.NewAbiFilterer(common.HexToAddress(wormholeCoreContract), backend)
	if err != nil {
		return fmt.Errorf("failed to bind client to wormhole contract: %w", err)
	}

	sink := make(chan *wormhole.AbiLogMessagePublished)

	// Subscribe to LogMessagePublished events.
	// The for loop ensures we continue to re-subscribe in case of a subscription error.
	for {
		var sub event.Subscription
		err := retry.Do(func() error {
			sub, err = wormholeBinding.WatchLogMessagePublished(
				&bind.WatchOpts{Context: ctx},
				sink,
				[]common.Address{common.HexToAddress(wormholeTransceiverContract)},
			)
			if err != nil {
				return err
			}
			return nil
		},
			retry.Context(ctx),
			retry.OnRetry(func(attempt uint, err error) {
				log.Error("retry: subscription", "attempt", attempt+1, "error", err)
			}),
		)
		if err != nil {
			return fmt.Errorf("failed to subscribe to `LogMessagePublished` events %w", err)
		}
		defer sub.Unsubscribe()

		log.Info("successfully subscribed to `LogMessagePublished` events")

		// listen and store data from events
		for {
			select {
			case <-ctx.Done():
				return nil
			case event := <-sink:
				log.Debug("observed `LogMessagePublished` event", "txHash", event.Raw.TxHash.String(), "sequence", event.Sequence)
				logMessagePublishedMap.Store(event.Raw.TxHash.String(), event.Sequence)
			case err := <-sub.Err():
				log.Error("subscription error. Re-subscribing.", "error", err)
				sub.Unsubscribe()
				// TODO: small historical query to catch up during re-subscription
				break
			}
		}
	}
}

// M0Listener listens for MTokenSent events
func M0Listener(
	ctx context.Context, log *slog.Logger,
	logMessagePublishedMap *LogMessagePublishedMap,
	ws *ethclient.Client,
	processingQueue chan *QueryData,
	wormholeSrcChainId uint64,
	hubPortalContract, wormholeTransceiverContract string,
) error {
	log = log.With(slog.String("listener", "m0"))

	backend := NewContractBackendWrapper(ws)
	binding, err := mportal.NewBindings(common.HexToAddress(hubPortalContract), backend)
	if err != nil {
		return fmt.Errorf("failed to bind client to mportal contract: %w", err)
	}

	sink := make(chan *mportal.BindingsMTokenSent)

	// Subscribe to MTokenSent events.
	// The for loop ensures we continue to re-subscribe in case of a subscription error.
	for {
		var sub event.Subscription
		err := retry.Do(func() error {
			sub, err = binding.WatchMTokenSent(
				&bind.WatchOpts{Context: ctx},
				sink,
				[]uint16{}, // TODO: update with destination chain ID
				nil, nil,
			)
			if err != nil {
				return err
			}
			return nil
		},
			retry.Context(ctx),
			retry.OnRetry(func(attempt uint, err error) {
				log.Error("retry: subscription", "attempt", attempt+1, "error", err)
			}),
		)
		if err != nil {
			return fmt.Errorf("failed to subscribe to `MTokenSent` events: %w", err)
		}

		defer sub.Unsubscribe()

		log.Info("successfully subscribed to `MTokenSent` events")

		// listen and process data from events
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
				err := retry.Do(func() error {
					seq, found = logMessagePublishedMap.GetSequence(txHash)
					if !found {
						return errors.New("not found")
					}
					return nil
				},
					retry.Context(ctx),
					retry.Attempts(3),
					retry.Delay(5*time.Millisecond),
					retry.OnRetry(func(attempt uint, _ error) {
						log.Debug("retry: logMessagePublished lookup", "attempt", attempt+1, "txHash", txHash)
					}),
				)
				if err != nil {
					log.Error("`logMessagePublished` sequence not found in correlation to `MTokenSent` event", "txHash", txHash)
				}
				if err == nil {
					log.Info("found correlating `logMessagePublished` event", "txHash", txHash, "sequence", seq)

					processingQueue <- &QueryData{
						WormHoleChainID: wormholeSrcChainId,
						Emitter:         wormholeTransceiverContract,
						Sequence:        seq,
						txHash:          txHash,
					}
					logMessagePublishedMap.Delete(event.Raw.TxHash.String())
				}

			case err := <-sub.Err():
				log.Error("subscription error. Re-subscribing.", "error", err)
				sub.Unsubscribe()
				// TODO: small historical query to catch up during re-subscription
				break
			}
		}
	}
}

// GetHistory queries historical data.
func GetHistory(
	ctx context.Context, log *slog.Logger,
	rpc *ethclient.Client,
	processingQueue chan *QueryData,
	startBlock int64, endBlock int64,
	wormholeSrcChainId uint64,
	hubPortalContract, wormholeCoreContract, wormholeTransceiverContract string,
) {
	from := big.NewInt(startBlock)
	var end *big.Int
	if endBlock != 0 {
		end = big.NewInt(endBlock)
	}

	log = log.With(slog.Int64("start-block", startBlock), slog.Int64("end-block", endBlock))
	log.Info("starting to query history")

	// load mPortal ABI and get MTokenSent function signature
	mPortalAbi, err := abi.JSON(strings.NewReader(mportal.BindingsMetaData.ABI))
	if err != nil {
		log.Error("unable to parse MTokenSent ABI when querying history", "error", err)
		return
	}
	mTokenSentFuncSig := mPortalAbi.Events["MTokenSent"].ID

	query := ethereum.FilterQuery{
		FromBlock: from,
		ToBlock:   end,
		Addresses: []common.Address{common.HexToAddress(hubPortalContract)},
		Topics:    [][]common.Hash{{mTokenSentFuncSig}},
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
	query.Addresses = []common.Address{common.HexToAddress(wormholeCoreContract)}
	query.Topics = [][]common.Hash{{logMessagePublishedFuncSig}, {common.HexToHash(wormholeTransceiverContract)}}

	logMessagePublishedLogs, err := rpc.FilterLogs(ctx, query)
	if err != nil {
		log.Error("unable to filter `logMessagePublished` logs when querying history", "error", err)
		return
	}

	event := struct {
		wormhole.AbiLogMessagePublished
	}{}

	var totalVaas int
	for _, mLog := range mTokenSentLogs {
		txHash := mLog.TxHash
		for _, lLog := range logMessagePublishedLogs {
			if txHash == lLog.TxHash {
				if err := wormholeAbi.UnpackIntoInterface(&event, "LogMessagePublished", lLog.Data); err != nil {
					log.Error("error unpacking wormhole abi into interface when querying history", "error", err)
				}
				log.Debug("vaa found during historical query", "seq", event.Sequence)
				processingQueue <- &QueryData{
					WormHoleChainID: wormholeSrcChainId,
					Emitter:         wormholeTransceiverContract,
					Sequence:        event.Sequence,
					txHash:          txHash.String(),
				}
				totalVaas += 1
			}
		}
	}

	defer log.Info("finished querying history", "vaas-found", totalVaas)
}
