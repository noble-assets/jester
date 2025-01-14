package appstate

import "github.com/spf13/cobra"

type Config struct {
	Log_level     string `toml:"log-level"`
	Log_style     string `toml:"log-style"`
	Testnet       bool   `toml:"testnet"`
	ServerAddress string `toml:"server-address"`

	Ethereum *Ethereum `toml:"ethereum"`
	Noble    *Noble    `toml:"noble"`
}

type Ethereum struct {
	WebsocketURL string `toml:"websocket-url"`
	RPCURL       string `toml:"rpc-url"`
}

type Noble struct {
	GRPCURL string `toml:"grpc-url"`
}

// AddConfigFlags adds flags that correspond to the config.toml file.
// The user has the option to use the flag, env var, or config file. Note that it
// takes precedent in that order.
//
// These are persistent flags.
// Ideally there is a flag for each config setting.
func AddConfigFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&eth_websocket, flagEthWebsocket, "w", "", "ethereum websocket")
	cmd.PersistentFlags().StringVarP(&eth_rpc, flagEthRPC, "r", "", "ethereum rpc")
	cmd.PersistentFlags().BoolVar(&testnet, flagTestnet, false, "use testnet configuration (contracts, chain ID's, ect.)")
	cmd.PersistentFlags().StringVar(&server_address, flagServerAddr, "localhost:9091", "gRPC server address")
	cmd.PersistentFlags().StringVar(&noble_grpc, flagNobleGRPC, "localhost:9090", "noble grpc address")
}
