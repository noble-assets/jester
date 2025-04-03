# ABI Generation

ABI files were generated from [`https://github.com/m0-foundation/m-portal`](https://github.com/m0-foundation/m-portal)

To Generate:

1) Inside the the `m-portal` repo, run `make build`
2) Run `abigen` EXAMPLE: `abigen --abi out/HubPortal.sol/HubPortal.abi.json --pkg mportal --out "HubPortal.go"`
