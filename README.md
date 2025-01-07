# Jester

Jester is a sidecar application designed to be run by the Noble validator set. It facilitates the implementation of the Noble Dollar, powered by M0.

## Installation

To install Jester, clone the repository and build the application:

```sh
git clone https://github.com/noble-assets/jester.git
cd jester
make build
# jesterd will be placed in the 'jester/build/' folder
```

Jester is also available via [releases](https://github.com/noble-assets/jester/releases) and [Docker](https://github.com/noble-assets/jester/pkgs/container/jester).

## Usage

Initialize config:

`jesterd config init`

By default, Jester's config will be placed in `$HOME/.jester/`.

An Ethereum `websocket_url` and `rpc_url` are required to run Jester.

By default, Jesters gRPC server address is `localhost:9091`. This endpoint should not be publicly exposed.
Querying the `query.v1.QueryService/GetVaas` should only be done by the Noble binary. Querying this endpoint clears cached VAA's.

## Flags and Configuration

Users have three options to configure Jester:

1. Flags: Use the --help flag to see available flags for specific commands.
2. Environment Variables: These must match either a config file setting or a flag. Must be UPPERCASE. If a flag uses "-", the environment variable equivalent is "_". There is one unique environment variable JESTER_HOME that mimics the --home flag.
3. Config File: See jester config init for more details.

Configuration precedence is as follows (from highest to lowest):

1. Flag
2. Environment Variable
3. Config File
