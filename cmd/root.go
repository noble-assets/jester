package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"jester.noble.xyz/appstate"
	"jester.noble.xyz/cmd/config"
)

const (
	appName         = "jesterd"
	defaultLogLevel = "info"
	defaultLogStyle = "pretty"
)

func NewRootCommand() *cobra.Command {
	a := &appstate.AppState{
		Viper: viper.New(),
	}

	rootCmd := &cobra.Command{
		Use:   appName,
		Short: "Jester is a 'sidecar' meant to run alongside the nobled binary.",
		Long: `Jester is a 'sidecar' meant to run alongside the nobled binary.
	
Jester is only necessary if you are also a validator.`,
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			// Inside persistent pre-run because this takes effect after flags are
			// parsed for both root and all child commands.
			a.ConfigureViper(cmd)
			a.LoadConfig()
			a.InitLogger()
		},
	}

	rootCmd.PersistentFlags().String(appstate.FlagHome, defaultHome(), "directory for config and data \n (optional env var = JESTER_HOME)")
	// manually bind "home" instead of using viper.AutomaticEnv
	if err := a.Viper.BindEnv(appstate.FlagHome, "JESTER_HOME"); err != nil {
		panic(err)
	}
	rootCmd.PersistentFlags().String(appstate.FlagLogLevel, defaultLogLevel, "log level format (info, debug, warn, error)")
	rootCmd.PersistentFlags().String(appstate.FlagLogStyle, defaultLogStyle, "log style format (text, json, pretty)")

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
