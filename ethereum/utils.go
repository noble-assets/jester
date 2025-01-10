package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractBackendWrapper struct {
	*ethclient.Client
}

// NewWebsocketContractBackendWrapper creates the backend for a websocket client.
// This is used to bind to an abi filter.
func NewWebsocketContractBackendWrapper(eth *Eth) *ContractBackendWrapper {
	eth.WebsocketClientMutex.Lock()
	defer eth.WebsocketClientMutex.Unlock()

	return &ContractBackendWrapper{Client: eth.WebsocketClient}
}
