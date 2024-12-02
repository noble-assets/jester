package appstate

type Config struct {
	Log_level string `toml:"log-level"`

	Ethereum_websocket string `toml:"ethereum-websocket"`
}
