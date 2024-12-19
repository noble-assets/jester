package appstate

type Config struct {
	Log_level     string `toml:"log_level"`
	Log_style     string `toml:"log_style"`
	Testnet       bool   `toml:"testnet"`
	ServerAddress string `toml:"server_address"`

	Ethereum *Ethereum `toml:"ethereum"`
}

type Ethereum struct {
	WebsocketURL string `toml:"websocket_url"`
	RPCURL       string `toml:"rpc_url"`
}
