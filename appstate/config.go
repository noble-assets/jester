package appstate

type Config struct {
	Log_level     string `toml:"log-level"`
	Log_style     string `toml:"log-style"`
	Testnet       bool   `toml:"testnet"`
	ServerAddress string `toml:"server-address"`

	Ethereum *Ethereum `toml:"ethereum"`
}

type Ethereum struct {
	WebsocketURL string `toml:"websocket-url"`
	RPCURL       string `toml:"rpc-url"`
}
