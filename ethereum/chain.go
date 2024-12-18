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
func InitializeEth(ctx context.Context, log *slog.Logger, websocketurl string, rpcurl string) (*Eth, error) {
	eth := newEth(websocketurl, rpcurl)
	if err := eth.initWebsocket(ctx, log); err != nil {
		return nil, err
	}
	if err := eth.initRPC(ctx, log); err != nil {
		return nil, err
	}
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
