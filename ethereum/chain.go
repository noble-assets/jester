package ethereum

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/ethclient"
	"jester.noble.xyz/metrics"
	"jester.noble.xyz/utils"
)

type Eth struct {
	Config               *Config
	WebsocketClient      *ethclient.Client
	WebsocketClientMutex sync.Mutex
	RPCClient            *ethclient.Client

	redial redial
}

type Config struct {
	WebsocketURL string
	RPCURL       string

	WormholeSrcChainId   uint16
	WormholeNobleChainID uint16
	WormholeApiUrl       string
	HubPortal            string
	WormholeCore         string
	WormholeTransceiver  string // LogMessagePublished topic sender
}

// redial is used to manage the redial state between one gRPC client
// and multiple websockets.
type redial struct {
	inProgressMutex sync.Mutex
	inProgress      bool
	done            chan error
	getHistory      chan struct{} // redial done, trigger historical lookup
}

type Overrides struct {
	WormholeSrcChainId   uint16
	WormholeNobleChainID uint16
	WormholeApiUrl       string
	HubPortal            string
	WormholeCore         string
	WormholeTransceiver  string
}

func newEth(websocketurl string, rpcurl string) *Eth {
	return &Eth{
		Config: &Config{
			WebsocketURL: websocketurl,
			RPCURL:       rpcurl,
		},
	}
}

// InitializeEth initializes Ethereum with a websocket and rpc client.
// The intent behind this is to have this command run during cobras `PreRunE` or
// `PersistentPreRunE`.
// The returned *Eth pointer should be added to the app state.
func InitializeEth(ctx context.Context, log *slog.Logger, websocketurl, rpcurl string, testnet bool, overrides Overrides) (*Eth, error) {
	eth := newEth(websocketurl, rpcurl)

	eth.setContracts(log, testnet, overrides)
	if err := eth.dialWebsocket(ctx, log); err != nil {
		return nil, err
	}
	if err := eth.dialRPC(ctx, log); err != nil {
		return nil, err
	}

	eth.redial.getHistory = make(chan struct{})

	return eth, nil
}

// dialRPC creates an Ethereum RPC client
func (e *Eth) dialRPC(ctx context.Context, log *slog.Logger) (err error) {
	err = retry.Do(func() error {
		e.RPCClient, err = ethclient.DialContext(ctx, e.Config.RPCURL)
		if err != nil {
			return err
		}
		return nil
	},
		retry.Context(ctx),
		retry.OnRetry(func(attempt uint, err error) {
			log.Warn("retrying to dial Ethereum RPC", "attempt", fmt.Sprintf("%d/%d", attempt+1, 10), "err", err)
		}))
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum RPC: %v", err)
	}

	log.Info("successfully dialed Ethereum RPC")

	return nil
}

// dialWebsocket creates an Ethereum websocket client
func (e *Eth) dialWebsocket(ctx context.Context, log *slog.Logger) (err error) {
	err = retry.Do(func() error {
		e.WebsocketClientMutex.Lock()
		defer e.WebsocketClientMutex.Unlock()

		e.WebsocketClient, err = ethclient.DialContext(ctx, e.Config.WebsocketURL)
		if err != nil {
			return err
		}
		return nil
	},
		retry.Context(ctx),
		retry.OnRetry(func(attempt uint, err error) {
			log.Warn("retrying to dial Ethereum websocket", "attempt", fmt.Sprintf("%d/%d", attempt+1, 10), "err", err)
		}))
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum WebSocket: %v", err)
	}

	log.Info("successfully dialed Ethereum websocket")

	return nil
}

// handleRedial handles the redial of the websocket client between multiple websocket subscriptions.
// Because the websocket client is shared between multiple subscriptions, this function
// is used to ensure that only one redial is in progress at a time.
func (e *Eth) handleRedial(ctx context.Context, log *slog.Logger, m *metrics.PrometheusMetrics) error {
	redial := &e.redial
	redial.inProgressMutex.Lock()

	// Another goroutine is already handling redial
	if redial.inProgress {
		log.Info("client redial already in progress")
		redialDone := redial.done // create own reference to avoid races
		redial.inProgressMutex.Unlock()
		err := <-redialDone
		errExists := false
		if err != nil {
			errExists = true
		}
		log.Info("received client redial complete", "error-exists", errExists)
		return err
	}

	// Mark redial as in progress and prepare a new signal channel
	redial.inProgress = true
	redial.done = make(chan error)
	redialDone := redial.done // create own reference to avoid races
	redial.inProgressMutex.Unlock()

	defer func() {
		redial.inProgressMutex.Lock()
		redial.inProgress = false
		close(redialDone)
		redial.inProgressMutex.Unlock()
	}()

	if m.Enabled {
		m.IncEthSubInterruptionCounter()
	}

	if err := e.dialWebsocket(ctx, log); err != nil {
		redialDone <- err
		return err
	}

	time.Sleep(2 * time.Second)     // Allow other redial signals to accumulate
	redial.getHistory <- struct{}{} // Trigger historical lookup to catch up on missed events
	return nil
}

// GetHistoricalOnRedial is used to catch up on any block data missed during an event
// subscription interruption. It is hardcoded to look back 50 blocks.
//
// It is meant to be run in a goroutine.m
func (e *Eth) GetHistoricalOnRedial(ctx context.Context, log *slog.Logger, processingQueue chan *utils.QueryData) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-e.redial.getHistory:
			lookback := uint64(50) // TODO: make lookback configurable
			log.Info(fmt.Sprintf("getting historical events for %d blocks", lookback))
			latest, err := e.RPCClient.BlockNumber(ctx)
			if err != nil {
				log.Error("failed to get latest block number", "error", err)
				continue
			}
			start := latest - lookback
			GetHistory(ctx, log, e, processingQueue, int64(start), 0)
		}
	}
}

func (e *Eth) CloseClients() {
	if e.WebsocketClient != nil {
		e.WebsocketClient.Close()
	}
	if e.RPCClient != nil {
		e.RPCClient.Close()
	}
}

// setContracts sets the contract addresses for the Ethereum network.
func (e *Eth) setContracts(log *slog.Logger, testnet bool, overrides Overrides) {
	switch testnet {
	case true:
		e.Config.WormholeSrcChainId = 10002
		e.Config.WormholeApiUrl = "https://api.testnet.wormscan.io/v1/signed_vaa"
		e.Config.HubPortal = "0x1B7aE194B20C555B9d999c835F74cDCE36A67a74"
		e.Config.WormholeCore = "0x4a8bc80Ed5a4067f1CCf107057b8270E0cC11A78"
		e.Config.WormholeTransceiver = "0x7B1bD7a6b4E61c2a123AC6BC2cbfC614437D0470"
	default:
		e.Config.WormholeSrcChainId = 2
		e.Config.WormholeApiUrl = ""      // TODO
		e.Config.HubPortal = ""           // TODO
		e.Config.WormholeCore = ""        // TODO
		e.Config.WormholeTransceiver = "" // TODO
	}

	e.Config.WormholeNobleChainID = 4009

	// Overrides
	if overrides.WormholeSrcChainId != 0 {
		log.Info("overriding wormhole source chain ID", "chainID", overrides.WormholeSrcChainId)
		e.Config.WormholeSrcChainId = overrides.WormholeSrcChainId
	}
	if overrides.WormholeNobleChainID != 0 {
		log.Info("overriding noble wormhole chain ID", "chainID", overrides.WormholeNobleChainID)
		e.Config.WormholeNobleChainID = overrides.WormholeNobleChainID
	}
	if overrides.WormholeApiUrl != "" {
		log.Info("overriding wormhole API URL", "url", overrides.WormholeApiUrl)
		e.Config.WormholeApiUrl = overrides.WormholeApiUrl
	}
	if overrides.HubPortal != "" {
		log.Info("overriding hub portal contract", "address", overrides.HubPortal)
		e.Config.HubPortal = overrides.HubPortal
	}
	if overrides.WormholeCore != "" {
		log.Info("overriding wormhole core contract", "address", overrides.WormholeCore)
		e.Config.WormholeCore = overrides.WormholeCore
	}
	if overrides.WormholeTransceiver != "" {
		log.Info("overriding wormhole transceiver contract", "address", overrides.WormholeTransceiver)
		e.Config.WormholeTransceiver = overrides.WormholeTransceiver
	}
}
