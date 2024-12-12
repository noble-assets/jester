//go:build tools

// This file uses the recommended method for tracking developer tools in a Go module.
// REF: https://go.dev/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "mvdan.cc/gofumpt"
)
