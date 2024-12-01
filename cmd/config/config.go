package config

import (
	"github.com/noble-assets/jester/configuration"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
func ConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configuration commands",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}

	configuration.AddConfigurationFlags(cmd)

	cmd.AddCommand(
		initCmd(),
	)

	return cmd
}
