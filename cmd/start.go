package cmd

import (
	"fmt"

	"github.com/noble-assets/jester/configuration"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "TODO",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("start called")

			v := viper.GetString("ethereum_websocket")
			fmt.Println(v)

		},
	}

	configuration.AddConfigurationFlags(cmd)

	return cmd
}
