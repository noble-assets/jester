package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractBackendWrapper struct {
	*ethclient.Client
}

// NewContractBackendWrapper creates a new ContractBackendWrapper.
// This is used to associate a websocket or rpc client an abi binding.
func NewContractBackendWrapper(client *ethclient.Client) *ContractBackendWrapper {
	return &ContractBackendWrapper{
		Client: client,
	}
}
