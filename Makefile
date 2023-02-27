BIN             := bin
SHELL           := /bin/bash
ZSH_COMPLETIONS := $(ZSH_CUSTOM)/plugins/completions

GOARCH ?= $(shell go env GOARCH)
GOBIN  ?= $(HOME)/.local/bin
GOOS   ?= $(shell go env GOOS)

ENV := env GOARCH=$(GOARCH) GOBIN=$(GOBIN) GOOS=$(GOOS)
GO  := $(ENV) go

build: | $(BIN)
	$(GO) build -o $(BIN) ./...
	$(ENV) $(SHELL) $(CURDIR)/scripts/rename.sh

clean:
	$(RM) -r $(BIN)

install:
	GOBIN=$(GOBIN) $(GO) install ./...
	pkg completion zsh > $(ZSH_COMPLETIONS)/_pkg
	utils completion zsh > $(ZSH_COMPLETIONS)/_utils

update:
	$(GO) get -u ./...

fmt:
	$(GO) fmt ./...
	$(GO) mod tidy

$(BIN):
	mkdir -p $@
