package appstate

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagHome         = "home"
	FlagLogLevel     = "log-level"
	FlagLogStyle     = "log-style"
	flagEthWebsocket = "websocket-url"
	flagEthRPC       = "rpc-url"
)

// if the flag is being added to multiple commands, you must use the flags that accepts pointers
// example `StringVar` instead of `String`
var (
	eth_websocket string
	eth_rpc       string
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
}
