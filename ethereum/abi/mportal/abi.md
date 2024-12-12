# ABI Generation

ABI files were generated from [`https://github.com/m0-foundation/m-portal`](https://github.com/m0-foundation/m-portal)

To Generate:

1) Inside the the `m-portal` repo, run `make build`
2) Run `abigen` EXAMPLE: `abigen --abi out/Address.sol/Address.abi.json --pkg abi --out "addresses.go"`

Note: If needed, this bash script can be copied to root and of `m-portal` repo and run to generate all bindings:

```sh
#!/bin/bash

OUT_DIR="./out"             # Directory where Foundry places compiled files
PACKAGE_NAME="bindings"     # Customize the package name
GENERATED_DIR="./generated" # Output directory for Go bindings

# Ensure output directory exists
mkdir -p "$GENERATED_DIR"

# Find all .abi.json files in the out directory
for ABI_FILE in $(find "$OUT_DIR" -name "*.abi.json"); do
    # Extract the contract name and path
    CONTRACT_NAME=$(basename "$ABI_FILE" .abi.json)
    CONTRACT_DIR=$(dirname "$ABI_FILE")
    BIN_FILE="$CONTRACT_DIR/$CONTRACT_NAME.json"

    # Check if the corresponding .bin exists
    if [ ! -f "$BIN_FILE" ]; then
        echo "Binary file for $CONTRACT_NAME not found. Skipping."
        continue
    fi

    # Generate Go bindings
    abigen --abi="$ABI_FILE" --pkg="$PACKAGE_NAME" --out="$GENERATED_DIR/$CONTRACT_NAME.go"

    echo "Generated Go bindings for $CONTRACT_NAME"
done
```
