package appstate

import (
	"github.com/spf13/cobra"
)

const (
	FlagHome                        = "home"
	FlagLogLevel                    = "log-level"
	FlagLogStyle                    = "log-style"
	flagTestnet                     = "testnet"
	flagEthWebsocket                = "ethereum.websocket-url"
	flagEthRPC                      = "ethereum.rpc-url"
	flagNobleGRPC                   = "noble.grpc-url"
	flagServerAddr                  = "server-address"
	FlagStartBlock                  = "start-block"
	FlagEndBlock                    = "end-block"
	FlagOverrideWormholeSrcChainId  = "override-wormhole-src-chain-id"
	FlagOverrideNobleChainID        = "override-noble-chain-id"
	FlagOverrideWormholeApiUrl      = "override-wormhole-api-url"
	FlagOverrideHubPortal           = "override-hub-portal"
	FlagOverrideWormholeTransceiver = "override-wormhole-transceiver"
	FlagOverrideWormholeCore        = "override-wormhole-core"
)

// if the flag is being added to multiple commands, you must use the flags that accepts pointers
// example `StringVar` instead of `String`
var (
	eth_websocket  string
	eth_rpc        string
	noble_grpc     string
	server_address string
	testnet        bool
)

// AddConfigurationFlags adds configuration related flags to a command.
// This helps streamline configuration with viper. The user has the option
// to use the flag, env var, or config file. Note that it takes precedent
// in that order.
//
// These are persistent flags.
// Ideally there is a flag for each config setting.
func AddConfigurationFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&eth_websocket, flagEthWebsocket, "w", "", "ethereum websocket")
	cmd.PersistentFlags().StringVarP(&eth_rpc, flagEthRPC, "r", "", "ethereum rpc")
	cmd.PersistentFlags().BoolVar(&testnet, flagTestnet, false, "use testnet configuration (contracts, chain ID's, ect.)")
	cmd.PersistentFlags().StringVar(&server_address, flagServerAddr, "localhost:9091", "gRPC server address")
	cmd.PersistentFlags().StringVar(&noble_grpc, flagNobleGRPC, "localhost:9090", "noble grpc address")
}
