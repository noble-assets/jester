package appstate

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/noble-assets/jester/ethereum"
	"github.com/phsym/console-slog"
	"github.com/spf13/viper"
)

// appState is the modifiable state of the application.
type AppState struct {
	Config *Config

	Log *slog.Logger

	// ethereum clients
	*ethereum.Eth
}

func (a *AppState) InitLogger() {
	var level slog.Level
	logLevel := strings.ToLower(viper.GetString(FlagLogLevel))

	switch logLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		fmt.Printf("\ninvalid log_level (%s); using 'info", logLevel)
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{Level: level}

	var logHandler slog.Handler
	logStyle := strings.ToLower(viper.GetString(FlagLogStyle))

	switch logStyle {
	case "text":
		logHandler = slog.NewTextHandler(os.Stderr, opts)
	case "json":
		logHandler = slog.NewJSONHandler(os.Stderr, opts)
	case "pretty":
		logHandler = console.NewHandler(os.Stderr, &console.HandlerOptions{Level: level})
	default:
		fmt.Printf("\ninvalid log_style (%s); using 'pretty'", logStyle)
		logHandler = console.NewHandler(os.Stderr, &console.HandlerOptions{Level: level})
	}

	a.Log = slog.New(logHandler)
}

// loadConfigFile reads configuration from file, env, OR args into a.Config
// This allows access to configuration via the appstate instead of using viper.
func (a *AppState) LoadConfig() {
	home := viper.GetString(FlagHome)

	viper.AddConfigPath(home)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.AutomaticEnv() // after reading in config, check for matching env vars
	viper.EnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.EnvKeyReplacer(strings.NewReplacer(".", "_"))

	a.Config = &Config{
		Log_level:     viper.GetString(FlagLogLevel),
		Log_style:     viper.GetString(FlagLogStyle),
		Testnet:       viper.GetBool(flagTestnet),
		ServerAddress: viper.GetString(flagServerAddr),
		Ethereum: &Ethereum{
			WebsocketURL: viper.GetString(flagEthWebsocket),
			RPCURL:       viper.GetString(flagEthRPC),
		},
	}
}
