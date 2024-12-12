package cmd

import (
	"os"
	"path/filepath"

	"github.com/noble-assets/jester/appstate"
	"github.com/noble-assets/jester/cmd/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	appName         = "jesterd"
	defaultLogLevel = "info"
	defaultLogStyle = "text"
)

func NewRootCommand() *cobra.Command {
	a := new(appstate.AppState)

	rootCmd := &cobra.Command{
		Use:   appName,
		Short: "Jester is a 'sidecar' meant to run alongside the nobled binary.",
		Long: `Jester is a 'sidecar' meant to run alongside the nobled binary.
	
Jester is only necessary if you are also a validator.`,
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			// Inside persistent pre-run because this takes effect after flags are parsed.
			a.LoadConfig()
			a.InitLogger()
		},
	}

	rootCmd.PersistentFlags().String(appstate.FlagHome, defaultHome(),
		"directory for config and data \n (optional env Var = JESTER_HOME)")
	if err := viper.BindPFlag(appstate.FlagHome, rootCmd.PersistentFlags().Lookup(appstate.FlagHome)); err != nil {
		panic(err)
	}
	// manually bind "home" instead of using viper.AutomaticEnv
	if err := viper.BindEnv(appstate.FlagHome, "JESTER_HOME"); err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().String(appstate.FlagLogLevel, defaultLogLevel,
		"log level format (info, debug, warn, error)")
	if err := viper.BindPFlag(appstate.FlagLogLevel, rootCmd.PersistentFlags().Lookup(appstate.FlagLogLevel)); err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().String(appstate.FlagLogStyle, defaultLogStyle,
		"log style format (text, json, console)")
	if err := viper.BindPFlag(appstate.FlagLogStyle, rootCmd.PersistentFlags().Lookup(appstate.FlagLogStyle)); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(
		config.ConfigCmd(a),
		startCmd(a),
		versionCmd(),
	)

	return rootCmd
}

func defaultHome() string {
	userHome, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(userHome, ".jester")
}
