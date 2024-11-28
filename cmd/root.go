package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/noble-assets/jester/cmd/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jester",
	Short: "Jester is a 'sidecar' meant to run alongside the nobled binary.",
	Long: `Jester is a 'sidecar' meant to run alongside the nobled binary.

Jester is only necessary if you are a validator.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Sub commands
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(versionCmd)

	// Set config path.
	userHome, err := os.UserHomeDir()
	cobra.CheckErr(err)
	home := filepath.Join(userHome, ".jester", "config.toml")

	// Global Flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", home, "config file path")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigFile(cfgFile)

	// TODO: Ensure this works! Currently not sure if viper.Get() is detecting flags arguments
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
}
