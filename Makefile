COMMIT := $(shell git log -1 --format='%H')
VERSION := $(shell echo $(shell git describe --tags --always --dirty --match "v*") | sed 's/^v//')


ldflags = -X github.com/noble-assets/jester/cmd.Version=$(VERSION) \
		  -X github.com/noble-assets/jester/cmd.Commit=$(COMMIT)

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

build:
	@echo "🤖 Building jester..."
	@go build -ldflags '$(ldflags)' -o "$(PWD)/build/jesterd"
	@echo "✅ Completed build!"

.PHONY: build
