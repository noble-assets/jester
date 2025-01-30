package appstate

const (
	FlagHome                        = "home"
	FlagLogLevel                    = "log-level"
	FlagLogStyle                    = "log-style"
	flagTestnet                     = "testnet"
	flagEthWebsocket                = "ethereum.websocket-url"
	flagEthRPC                      = "ethereum.rpc-url"
	flagServerAddr                  = "server-address"
	FlagStartBlock                  = "start-block"
	FlagEndBlock                    = "end-block"
	FlagOverrideWormholeSrcChainId  = "override-wormhole-src-chain-id"
	FlagOverrideNobleChainID        = "override-noble-chain-id"
	FlagOverrideWormholeApiUrl      = "override-wormhole-api-url"
	FlagOverrideHubPortal           = "override-hub-portal"
	FlagOverrideWormholeTransceiver = "override-wormhole-transceiver"
	FlagOverrideWormholeCore        = "override-wormhole-core"
	flagMetricsEnabled              = "metrics.enabled"
	flagMetricsAddress              = "metrics.address"
	FlagHushLogo                    = "hush-logo"
)

// if the flag is being added to multiple commands, you must use the flags that accepts pointers
// example `StringVar` instead of `String`
var (
	eth_websocket   string
	eth_rpc         string
	server_address  string
	testnet         bool
	metrics_enabled bool
	metrics_address string
)
