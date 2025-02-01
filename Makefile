COMMIT := $(shell git log -1 --format='%H')
VERSION := $(shell echo $(shell git describe --tags --always --dirty --match "v*"))

ldflags = -X jester.noble.xyz/cmd.Version=$(VERSION) \
		  -X jester.noble.xyz/cmd.Commit=$(COMMIT)

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

###############################################################################
###                                  Build                                  ###
###############################################################################

install:
	@echo "🤖 Installing Jester..."
	@go install -mod=readonly -ldflags '$(ldflags)' ./cmd/jesterd
	@echo "✅ Completed install!"

build:
	@echo "🤖 Building jester..."
	@go build -ldflags '$(ldflags)' -o "$(PWD)/build/" ./cmd/jesterd
	@echo "✅ Completed build!"

build-race:
	@echo "🤖 Building jester with race detection"
	@go build -ldflags '$(ldflags)' -o "$(PWD)/build/" ./cmd/jesterd -race
	@echo "✅ Completed build!"

###############################################################################
###                                 Tooling                                 ###
###############################################################################

gofumpt_cmd=mvdan.cc/gofumpt
golangci_lint_cmd=github.com/golangci/golangci-lint/cmd/golangci-lint

FILES := $(shell find . -name "*.go" -not -path "./ethereum/abi/*" -not -name "*.pb.go" -not -name "*.connect.go")
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
###                                Protobuf                                 ###
###############################################################################

BUF_VERSION=1.49

proto-all: proto-format proto-lint proto-gen

proto-format:
	@echo "🤖 Running protobuf formatter..."
	@docker run --rm --volume "$(PWD)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) format --diff --write
	@echo "✅ Completed protobuf formatting!"

proto-gen:
	@echo "🤖 Generating code from protobuf..."
	@docker run --rm --volume "$(PWD)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) generate
	@cp -r jester.noble.xyz/* .
	@rm -rf jester.noble.xyz
	@echo "✅ Completed code generation!"

proto-lint:
	@echo "🤖 Running protobuf linter..."
	@docker run --rm --volume "$(PWD)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) lint
	@echo "✅ Completed protobuf linting!"

###############################################################################

.PHONY: build license format lint proto-format proto-gen proto-lint
