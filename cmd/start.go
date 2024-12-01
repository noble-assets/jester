package cmd

import (
	"fmt"

	"github.com/noble-assets/jester/appstate"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
func startCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "TODO",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("start called")

			fmt.Println(a.Config.Ethereum_websocket)
		},
	}

	appstate.AddConfigurationFlags(cmd)

	return cmd
}
