package appstate

type Config struct {
	Log_level string `toml:"log-level"`
	Log_style string `toml:"log-style"`

	Ethereum *Ethereum
}

type Ethereum struct {
	WebsocketURL string `toml:"websocket-url"`
	RPCURL       string `toml:"rpc-url"`
}
