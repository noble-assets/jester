package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/noble-assets/jester/appstate"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
func initCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the config file",
		Long: `Initialize the config file. Optionally use flags to 
populate values.`,

		RunE: func(_ *cobra.Command, _ []string) error {
			configPath := filepath.Join(viper.GetString(appstate.FlagHome), ".jester", "config.toml")

			// programmatically overwriting config file does not treat viper keys ase expected
			// instead, notify user and have them delete config manually
			if fileExists(configPath) {
				return errors.New(fmt.Sprintf("file '%s' already exists", configPath))
			}

			// create config file
			dir := filepath.Dir(configPath)
			if err := os.MkdirAll(dir, 0o755); err != nil {
				return fmt.Errorf("failed to create directories: %w", err)
			}
			file, err := os.Create(configPath)
			if err != nil {
				return err
			}
			defer file.Close()

			// ensure valid toml
			encoder := toml.NewEncoder(file)
			err = encoder.Encode(a.Config)
			if err != nil {
				return fmt.Errorf("invalid config created: %v", err)
			}
			fmt.Printf("Config created at: '%s'\n", configPath)

			return nil
		},
	}

	return cmd
}

// fileExists checks if a file exists
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
