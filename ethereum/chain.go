package ethereum

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Eth struct {
	Config             *Config
	EthWebsocketClient *ethclient.Client
	EthRPCClient       *ethclient.Client
}

type Config struct {
	WebsocketURL string
	RPCURL       string

	WormholeSrcChainId  uint64
	WormholeApiUrl      string
	HubPortal           string
	WormholeCore        string
	WormholeTransceiver string // LogMessagePublished topic sender
}

type Overrides struct {
	WormholeSrcChainId  uint64
	WormholeApiUrl      string
	HubPortal           string
	WormholeCore        string
	WormholeTransceiver string
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
	if err := eth.initWebsocket(ctx, log); err != nil {
		return nil, err
	}
	if err := eth.initRPC(ctx, log); err != nil {
		return nil, err
	}

	eth.setContracts(log, testnet, overrides)

	return eth, nil
}

// initWebsocket creates an Ethereum websocket client
// If a WS client already exists, nothing is done.
func (e *Eth) initWebsocket(ctx context.Context, log *slog.Logger) (err error) {
	if e.EthWebsocketClient != nil {
		fmt.Println("eth websocket client already inialized")
		return nil
	}

	e.EthWebsocketClient, err = ethclient.DialContext(ctx, e.Config.WebsocketURL)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum WebSocket: %v", err)
	}
	log.Info("successfully dialed Ethereum websocket")

	return nil
}

// initRPC creates an Ethereum RPC client
// If a WS client already exists, nothing is done.
func (e *Eth) initRPC(ctx context.Context, log *slog.Logger) (err error) {
	if e.EthRPCClient != nil {
		fmt.Println("eth RPC client already inialized")
		return nil
	}

	e.EthRPCClient, err = ethclient.DialContext(ctx, e.Config.RPCURL)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum RPC: %v", err)
	}
	log.Info("successfully dialed Ethereum RPC")

	return nil
}

func (e *Eth) CloseClients() {
	if e.EthWebsocketClient != nil {
		e.EthWebsocketClient.Close()
	}
	if e.EthRPCClient != nil {
		e.EthRPCClient.Close()
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

	// Overrides
	if overrides.WormholeSrcChainId != 0 {
		log.Info("overriding wormhole source chain ID", "chainID", overrides.WormholeSrcChainId)
		e.Config.WormholeSrcChainId = overrides.WormholeSrcChainId
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
