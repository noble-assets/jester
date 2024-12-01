package appstate

type Config struct {
	Log_level string `toml:"log_level"`

	Ethereum_websocket string `toml:"ethereum_websocket"`
}
