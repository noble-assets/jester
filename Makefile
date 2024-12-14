APPNAME := jesterd
COMMIT := $(shell git log -1 --format='%H')
VERSION := $(shell echo $(shell git describe --tags --always --dirty --match "v*") | sed 's/^v//')

ldflags = -X github.com/noble-assets/jester/cmd.Version=$(VERSION) \
		  -X github.com/noble-assets/jester/cmd.Commit=$(COMMIT)

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

###############################################################################
###                                  Build                                  ###
###############################################################################

build:
	@echo "🤖 Building jester..."
	@go build -ldflags '$(ldflags)' -o "$(PWD)/build/$(APPNAME)"
	@echo "✅ Completed build!"

build-race:
	@echo "🤖 Building jester..."
	@go build -ldflags '$(ldflags)' -o "$(PWD)/build/$(APPNAME)" -race
	@echo "✅ Completed build!"

###############################################################################
###                                 Tooling                                 ###
###############################################################################

gofumpt_cmd=mvdan.cc/gofumpt
golangci_lint_cmd=github.com/golangci/golangci-lint/cmd/golangci-lint

FILES := $(shell find $(shell go list -f '{{.Dir}}' ./...) -name "*.go" -a -not -name "*.pb.go" -a -not -name "*.pb.gw.go" -a -not -name "*.pulsar.go" | sed "s|$(shell pwd)/||g")
license:
	@go-license --config .github/license.yml $(FILES)

format:
	@echo "🤖 Running formatter..."
	@go run $(gofumpt_cmd) -l -w .
	@echo "✅ Completed formatting!"

lint:
	@echo "🤖 Running linter..."
	@go run $(golangci_lint_cmd) run --timeout=10m
	@echo "✅ Completed linting!"

###############################################################################

.PHONY: build license format lint