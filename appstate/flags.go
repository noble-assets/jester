package appstate

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagHome              = "home"
	flagEthereumWebsocket = "ethereum_websocket"
)

// AddConfigurationFlags adds configuration related flags to a command.
// This helps streamline configuration with viper. The user has the option
// to use the flag, env var, or config file. Note that it takes precedent
// in that order.
//
// Ideally there is a flag for each config setting.
func AddConfigurationFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP(flagEthereumWebsocket, "w", "", "ethereum websocket")
	if err := viper.BindPFlag(flagEthereumWebsocket, cmd.PersistentFlags().Lookup(flagEthereumWebsocket)); err != nil {
		panic(err)
	}
}
