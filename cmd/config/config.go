package config

import (
	"github.com/spf13/cobra"
	"jester.noble.xyz/appstate"
)

// configCmd represents the config command
func ConfigCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configuration commands",
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
	}

	appstate.AddConfigurationFlags(cmd)

	cmd.AddCommand(
		initCmd(a),
	)

	return cmd
}
