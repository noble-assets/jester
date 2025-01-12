package appstate

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/phsym/console-slog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"jester.noble.xyz/ethereum"
	"jester.noble.xyz/noble"
)

// appState is the modifiable state of the application.
type AppState struct {
	Viper  *viper.Viper
	Config *Config

	Log *slog.Logger

	// ethereum clients
	*ethereum.Eth

	*noble.Noble
}

func (a *AppState) InitLogger() {
	viper := a.Viper
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

func (a *AppState) ConfigureViper(cmd *cobra.Command) {
	viper := a.Viper

	// bindAllFlags binds all flags to viper (both local and persistent).
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := viper.BindPFlag(f.Name, cmd.Flags().Lookup(f.Name)); err != nil {
			panic(err)
		}
	})

	home := viper.GetString(FlagHome)

	viper.AddConfigPath(home)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.AutomaticEnv() // after reading in config, check for matching env vars
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
}

// loadConfigFile reads configuration from file, env, OR args into a.Config
// This allows access to configuration via the appstate instead of using viper.
func (a *AppState) LoadConfig() {
	viper := a.Viper

	a.Config = &Config{
		Log_level:     viper.GetString(FlagLogLevel),
		Log_style:     viper.GetString(FlagLogStyle),
		Testnet:       viper.GetBool(flagTestnet),
		ServerAddress: viper.GetString(flagServerAddr),
		Ethereum: &Ethereum{
			WebsocketURL: viper.GetString(flagEthWebsocket),
			RPCURL:       viper.GetString(flagEthRPC),
		},
		Noble: &Noble{
			GRPCURL: viper.GetString(flagNobleGRPC),
		},
	}
}
