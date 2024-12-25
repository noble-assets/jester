package appstate

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagHome                        = "home"
	FlagLogLevel                    = "log_level"
	FlagLogStyle                    = "log_style"
	flagTestnet                     = "testnet"
	flagEthWebsocket                = "ethereum.websocket_url"
	flagEthRPC                      = "ethereum.rpc_url"
	flagServerAddr                  = "server_address"
	FlagStartBlock                  = "start_block"
	FlagEndBlock                    = "end_block"
	FlagOverrideWormholeSrcChainId  = "override_wormhole_src_chain_id"
	FlagOverrideNobleChainID        = "override_noble_chain_id"
	FlagOverrideWormholeApiUrl      = "override_wormhole_api_url"
	FlagOverrideHubPortal           = "override-hub-portal"
	FlagOverrideWormholeTransceiver = "override-wormhole-transceiver"
	FlagOverrideWormholeCore        = "override-wormhole-core"
)

// if the flag is being added to multiple commands, you must use the flags that accepts pointers
// example `StringVar` instead of `String`
var (
	eth_websocket  string
	eth_rpc        string
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
	if err := viper.BindPFlag(flagEthWebsocket, cmd.PersistentFlags().Lookup(flagEthWebsocket)); err != nil {
		panic(err)
	}

	cmd.PersistentFlags().StringVarP(&eth_rpc, flagEthRPC, "r", "", "ethereum rpc")
	if err := viper.BindPFlag(flagEthRPC, cmd.PersistentFlags().Lookup(flagEthRPC)); err != nil {
		panic(err)
	}

	cmd.PersistentFlags().BoolVar(&testnet, flagTestnet, false, "use testnet Ethereum contracts")
	if err := viper.BindPFlag(flagTestnet, cmd.PersistentFlags().Lookup(flagTestnet)); err != nil {
		panic(err)
	}

	cmd.PersistentFlags().StringVar(&server_address, flagServerAddr, "localhost:9091", "gRPC server address")
	if err := viper.BindPFlag(flagServerAddr, cmd.PersistentFlags().Lookup(flagServerAddr)); err != nil {
		panic(err)
	}
}
