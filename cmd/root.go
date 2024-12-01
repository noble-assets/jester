package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/noble-assets/jester/cmd/config"
	"github.com/noble-assets/jester/configuration"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "jester",
		Short: "Jester is a 'sidecar' meant to run alongside the nobled binary.",
		Long: `Jester is a 'sidecar' meant to run alongside the nobled binary.
	
Jester is only necessary if you are also a validator.`,
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			// Inside persistent pre-run because this takes effect after flags are parsed.
			initConfig()
		},
	}

	rootCmd.PersistentFlags().String(configuration.FlagHome, defaultHome(),
		"directory for config and data \n (optinoal env Var = JESTER_HOME)")
	if err := viper.BindPFlag(configuration.FlagHome, rootCmd.PersistentFlags().Lookup(configuration.FlagHome)); err != nil {
		panic(err)
	}
	// manually bind "home" instead of using viper.AutomaticEnv
	viper.BindEnv(configuration.FlagHome, "JESTER_HOME")

	rootCmd.AddCommand(
		config.ConfigCmd(),
		startCmd(),
		versionCmd(),
	)

	return rootCmd
}

func defaultHome() string {
	userHome, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(userHome, ".jester")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home := viper.GetString(configuration.FlagHome)

	// Search config in home directory with name ".cobra" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.AutomaticEnv() // after reading in config, check for matching env vars

	return
}
