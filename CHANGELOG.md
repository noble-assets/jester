# CHANGELOG

## v1.0.2

*Apr 24, 2025*

### FEATURES

- Add more Prometheus metrics and a no-op developer mode to assist in running Jester when you are not a validator. ([#16](https://github.com/noble-assets/jester/pull/16))
- Increase the default max attempts to fetch a VAA and add a `--override-fetch-vaa-attempts` flag. ([#17](https://github.com/noble-assets/jester/pull/17))

### FIX

- Add context to network requests and catch all timeout errors in retry logic.([#25](https://github.com/noble-assets/jester/pull/25))
- Fix an issue where Jester could crash if the user disables the metrics server. ([#26](https://github.com/noble-assets/jester/pull/26))

## v1.0.1

*Mar 14, 2025*

### BUG FIXES

- Correctly log `blockNumber` in debug logs. ([#14](https://github.com/noble-assets/jester/pull/14))
- Catch context `DeadlineExceeded` error and retry Wormhole query. ([#18](https://github.com/noble-assets/jester/pull/18))

### IMPROVEMENTS

- Improve logging when redialing clients. ([#15](https://github.com/noble-assets/jester/pull/15))

## v1.0.0

*Mar 3, 2025*

Introduces the first production-ready version of Jester.

No changes have been made since the release candidate v1.0.0-rc.1.

## v1.0.0-rc.1

*Feb 26, 2025*

### BUG FIXES

- Update ABI to be compatible with the latest design of the M Hub Portal. ([#12](https://github.com/noble-assets/jester/pull/12))

### FEATURE

- Add default mainnet config values. ([#3](https://github.com/noble-assets/jester/pull/3))

## v1.0.0-rc.0

*Feb 19, 2025*

### BREAKING CHANGES

- Combine both custom and default Prometheus metrics into one endpoint. ([#5](https://github.com/noble-assets/jester/issues/5))

### DEPENDENCIES

- Bump Go version to [`v1.24`](https://go.dev/doc/go1.24). ([#8](https://github.com/noble-assets/jester/issues/8))

### FEATURE

- Add Dockerfile. ([#6](https://github.com/noble-assets/jester/issues/6))

### REFACTOR

- Update default testnet config values. ([#4](https://github.com/noble-assets/jester/pull/4))

## v1.0.0-alpha.0

*Jan 31, 2025*

Introduces the initial release of the Jester binary.

