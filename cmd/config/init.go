package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/noble-assets/jester/config"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the config file",
	Long: `Initialize the config file. Optionally use flags to 
populate values.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}

		if fileExists(configPath) {
			// Ask the user if they want to overwrite
			fmt.Printf("File '%s' already exists. Do you want to overwrite it? (y/N): ", configPath)
			reader := bufio.NewReader(os.Stdin)
			response, _ := reader.ReadString('\n')
			response = strings.TrimSpace(strings.ToLower(response))

			if response != "y" {
				fmt.Println("Aborting!")
				return nil
			}
			fmt.Println("Overwriting config...")
		}

		config := config.Config{}

		// parse flags
		arg_ethereum_websocket, err := cmd.Flags().GetString("ethereum_websocket")
		if err != nil {
			return err
		}
		if arg_ethereum_websocket != "" {
			config.Ethereum_websocket = arg_ethereum_websocket
		}

		// create config file
		dir := filepath.Dir(configPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directories: %w", err)
		}
		file, err := os.Create(configPath)
		if err != nil {
			return err
		}
		defer file.Close()

		//ensure valid toml
		encoder := toml.NewEncoder(file)
		err = encoder.Encode(config)
		if err != nil {
			return fmt.Errorf("invalid config created: %v", err)
		}
		fmt.Printf("Config created at: '%s'\n", configPath)

		return nil
	},
}

// fileExists checks if a file exists
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func init() {
	config.AddConfigurationFlags(initCmd)
}
