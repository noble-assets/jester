package appstate

import (
	"fmt"

	"github.com/spf13/viper"
)

// appState is the modifiable state of the application.
type AppState struct {
	Config *Config
}

// loadConfigFile reads configuration from file, env, OR args into a.Config
// This allows access to configuration via the appstate instead of using viper.
func (a *AppState) LoadConfig() {
	home := viper.GetString(FlagHome)

	// Search config in home directory with name ".cobra" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.AutomaticEnv() // after reading in config, check for matching env vars

	a.Config = &Config{
		Ethereum_websocket: viper.GetString(flagEthereumWebsocket),
	}

	return
}
