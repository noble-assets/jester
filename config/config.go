package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Ethereum_websocket string `toml:"ethereum_websocket"`
}

// AddConfigurationFlags help streamline configuration with viper.
// The user has the option to use the config file OR pass in
// a flag arguement if they prefer to not have a config file.
//
// Ideally there is a flag for each config setting.
func AddConfigurationFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("ethereum_websocket", "w", "", "ethereum websocket")
	viper.BindPFlag("ethereum_websocket", cmd.Flags().Lookup("ethereum_websocket"))
}
