package appstate

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagHome              = "home"
	FlagLogLevel          = "log-level"
	flagEthereumWebsocket = "ethereum-websocket"
)

// if the flag is being added to multiple commands, you must use the flags that accept points
// example `StringVar` instead of `String`
var ethereum_websocket string

// AddConfigurationFlags adds configuration related flags to a command.
// This helps streamline configuration with viper. The user has the option
// to use the flag, env var, or config file. Note that it takes precedent
// in that order.
//
// These are persistent flags.
// Ideally there is a flag for each config setting.
func AddConfigurationFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&ethereum_websocket, flagEthereumWebsocket, "w", "", "ethereum websocket")
	if err := viper.BindPFlag(flagEthereumWebsocket, cmd.Flags().Lookup(flagEthereumWebsocket)); err != nil {
		panic(err)
	}
	return
}
