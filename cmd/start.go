package cmd

import (
	"github.com/noble-assets/jester/appstate"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
func startCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "TODO",

		Run: func(cmd *cobra.Command, args []string) {
			a.Log.Info("Start Called!")

			a.Log.Error(a.Config.Ethereum_websocket)
			a.Log.Info(a.Config.Log_level)
		},
	}

	appstate.AddConfigurationFlags(cmd)

	return cmd
}
