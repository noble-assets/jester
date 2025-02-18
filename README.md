# Jester

![GitHub Release](https://img.shields.io/github/v/release/noble-assets/jester?filter=v*)
![GitHub Release](https://img.shields.io/github/v/release/noble-assets/jester?filter=api%2Fv*)
![Build Workflow](https://github.com/noble-assets/jester/actions/workflows/release.yaml/badge.svg?event=push)
[![Go Reference](https://pkg.go.dev/badge/jester.noble.zyx.svg)](https://pkg.go.dev/jester.noble.xyz)

Jester is a sidecar application designed to be run by the Noble validator set. It facilitates the implementation of the Noble Dollar, powered by M0.

> [!NOTE]
> Jester is not needed by node operators. Only validators need to run Jester.

## What does Jester do?

Jester acts as a relayer listening for `$USDN` transfers between Ethereum and Noble.

As events are observed, Jester accumulates Wormhole VAA's (Verified Action Approvals).
These VAA's are handed off to Noble validators and submitted on-chain.

## Installation 💾

To install from source, clone the repository and install:

```sh
git clone https://github.com/noble-assets/jester.git
cd jester
make install
```

Jester is also available via [releases](https://github.com/noble-assets/jester/releases) and [Docker](https://github.com/noble-assets/jester/pkgs/container/jester).

To ensure is Jester is properly installed, try printing out version info:

```sh
jesterd version
```

## Setup ⚙

Initialize config:

```sh
jesterd config init
```

By default, Jester's config will be placed in `$HOME/.jester/`

<details>
<summary>Example Config</summary>

```toml
# log level format (info, debug, warn, error)
log-level = "debug"
# log style format (text, json, pretty)
log-style = "pretty"
testnet = false
# jester's gRPC server
server-address = "localhost:9091"

[ethereum]
  websocket-url = "wss://MY-ENDPOINT"
  rpc-url = "https://MY-ENDPOINT"

# Prometheus Metrics
[metrics]
  enabled = true
  address = "localhost:2112"
```

</details>

Jesters config is simple and mostly auto generated via the `init` command.  Please ensure all values are filled out.

You will need a reliable Ethereum `websocket-url` and Ethereum `rpc-url`.

Make sure you toggle the `testnet` boolean accordingly.

By default, Jesters gRPC server address is `localhost:9091`. This endpoint should not be publicly exposed and only accessed by the Noble binary.

<details>
<summary><h4>Are you are a validator updating from a Noble version < <code>v9</code>?</h4></summary>

You will need to add the following to Nobles `app.toml` file. By default this file lives at `$HOME/.noble/config/app.toml`

```toml
###############################################################################
###                             Jester (sidecar)                            ###
###############################################################################

[jester]

# Jester's gRPC server address. 
# This should not conflict with the CometBFT gRPC server.
grpc-address = "localhost:9091"
```

> New versions of Noble will automatically have this field when initialized.

</details>

Make sure that Jesters gRPC server matches the configuration set in Noble's `app.toml` (`$HOME/.noble/config/app.toml`).
You do not need need to worry about this if you are using defaults.

## Run 🃏

To run:

`jesterd start`

## Prometheus

Jester includes a Prometheus server to monitor its status. We **HIGHLY** encourage all operators to scrape the `up` metric to ensure that Jester is up and running.

By default, the Prometheus server listens on `localhost:2112`.

Both [custom metrics](./docs/prometheus.md#custom-metrics) and default Golang metrics are served at `/metrics`.

More details about the Prometheus server are available [here](./docs/prometheus.md).

> [!IMPORTANT]
> The only way Jester is expected to quit is if the Ethereum Websocket becomes disconnect.
> Jester will make several attempts to reconnect. If all attempts fails, Jester quits.
> Scraping the `up` metric will allow operators to be notified if this happens.

## Flags and Configuration

Users have three options to configure Jester:

1. Flags: Use the `--help` flag to see available flags per command.
2. Environment Variables: These must match either a config file setting or a flag. Must be UPPERCASE. If a flag uses "-", the ENV var equivalent is "_". There is one unique variable: `JESTER_HOME` that mimics the --home flag.
3. Config File

Note that configuration precedence takes place in the above order. Flags are prioritized, then ENV variables then the config file.

> TIP: Use flags when running `jesterd config init` to automatically fill out the config file.

## More Questions?

Please visit our [FAQ](./docs/faq.md).
