package ethereum

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
func WormholeListener(ctx context.Context, logMessagePublishedMap *LogMessagePublishedMap, log *slog.Logger, ws *ethclient.Client) error {
	log = log.With(slog.String("listener", "wormhole"))
	// TODO: Add to config
	wormholeSepoliaContract := "0x4a8bc80Ed5a4067f1CCf107057b8270E0cC11A78"

	backend := NewContractBackendWrapper(ws)
	opts := &bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}

	wormholeBinding, err := wormhole.NewAbiFilterer(common.HexToAddress(wormholeSepoliaContract), backend)
	if err != nil {
		return fmt.Errorf("failed to bind client to wormhole contract: %w", err)
	}

	sink := make(chan *wormhole.AbiLogMessagePublished)
	// TODO: How can we derive  this sender. Will it always be the same?
	logMessagePublishedSender := common.HexToAddress("0x7B1bD7a6b4E61c2a123AC6BC2cbfC614437D0470")
	sub, err := wormholeBinding.WatchLogMessagePublished(opts, sink, []common.Address{logMessagePublishedSender})
	if err != nil {
		return fmt.Errorf("failed to subscribe to `LogMessagePublished` events %w", err)
	}
	defer sub.Unsubscribe()

	fmt.Println("Successfully subscribed to `LogMessagePublished` events")

	for {
		select {
		case <-ctx.Done():
			return nil
		case event := <-sink:
			log.Debug("observed `LogMessagePublished` event", "txHash", event.Raw.TxHash.String(), "sequence", event.Sequence)
			logMessagePublishedMap.Store(event.Raw.TxHash.String(), event.Sequence)
		case err := <-sub.Err():
			// TODO: handle subscription interruption
			log.Error("Subscription error", "error", err)
		}
	}
}

// M0Listener listens for MTokenSent events
func M0Listener(ctx context.Context, logMessagePublishedMap *LogMessagePublishedMap, log *slog.Logger, ws *ethclient.Client, processingQueue chan *QueryData) error {
	log = log.With(slog.String("listener", "m0"))
	// Todo add to config
	// example tx hash on sepolia: "0xc2594725d7262d99b9ee523e70457cf614292ac9e8467d6fc92a79cca92a735c"
	m0SepoliaContract := "0x1B7aE194B20C555B9d999c835F74cDCE36A67a74"

	backend := NewContractBackendWrapper(ws)
	opts := &bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}

	binding, err := mportal.NewBindings(common.HexToAddress(m0SepoliaContract), backend)
	if err != nil {
		return fmt.Errorf("failed to bind client to mportal contract: %w", err)
	}

	sink := make(chan *mportal.BindingsMTokenSent)
	sub, err := binding.WatchMTokenSent(opts, sink, []uint16{}, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to subscribe to `MTokenSent` events: %w", err)
	}
	defer sub.Unsubscribe()

	fmt.Println("Successfully subscribed to `MTokenSent` events")

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
				retry.Attempts(3),
				retry.Delay(5*time.Millisecond),
				retry.OnRetry(func(attempt uint, _ error) {
					log.Debug("retry attempt: logMessagePublished lookup", "attempt", attempt+1, "txHash", txHash)
				}),
			)
			if err != nil {
				log.Error("`logMessagePublished` sequence not found in correlation to `MTokenSent` event", "txHash", txHash)
			}
			if err == nil {
				log.Info("found correlating `logMessagePublished` event", "txHash", txHash, "sequence", seq)

				processingQueue <- &QueryData{
					WormHoleChainID: 10002,                                        // TODO
					Emitter:         "0x7B1bD7a6b4E61c2a123AC6BC2cbfC614437D0470", // TODO
					Sequence:        seq,
					txHash:          txHash,
				}
				logMessagePublishedMap.Delete(event.Raw.TxHash.String())
			}

		case err := <-sub.Err():
			// TODO: handle subscription interruption
			log.Error("Subscription error", "error", err)
		}
	}
}
