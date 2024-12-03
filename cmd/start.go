package cmd

import (
	"github.com/noble-assets/jester/appstate"
	"github.com/noble-assets/jester/ethereum"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
func startCmd(a *appstate.AppState) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "TODO",
		PreRunE: func(cmd *cobra.Command, _ []string) (err error) {
			a.Eth, err = ethereum.InitializeEth(a.Config.Ethereum.WebsocketURL, a.Config.Ethereum.RPCURL, cmd.Context())
			if err != nil {
				return err
			}
			defer a.Eth.CloseClients()

			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
	}

	appstate.AddConfigurationFlags(cmd)

	return cmd
}
