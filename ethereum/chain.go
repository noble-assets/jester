package ethereum

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Eth struct {
	Config             *Config
	EthWebsocketClient *ethclient.Client
	EthRPCClient       *ethclient.Client // rpc.Client wrapped in ethclient.NewClient
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

func InitializeEth(websocketurl string, rpcurl string, ctx context.Context) (*Eth, error) {
	eth := newEth(websocketurl, rpcurl)
	if err := eth.initWebsocket(ctx); err != nil {
		return nil, err
	}
	// TODO init RPC client

	return eth, nil
}

func (e *Eth) WebsocketURL() string {
	return e.Config.WebsocketURL
}

func (e *Eth) RpcUrl() string {
	return e.Config.RPCURL
}

// initWebsocket creates an Ethereum websocket client
// If a WS client already exists, nothing is done.
func (e *Eth) initWebsocket(ctx context.Context) (err error) {
	if e.EthWebsocketClient != nil {
		fmt.Println("eth websocket client already inialized")
		return nil
	}

	e.EthWebsocketClient, err = ethclient.DialContext(ctx, e.Config.WebsocketURL)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum WebSocket: %v", err)
	}
	fmt.Println("Successfully dialed Ethereum websocket")

	return nil
}

func (e *Eth) CloseClients() {
	if e.EthWebsocketClient != nil {
		e.EthWebsocketClient.Close()
	}
}
