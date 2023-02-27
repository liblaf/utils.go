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
	$(RM) docs/pkg/pkg*.md
	$(RM) docs/utils/utils*.md

install:
	GOBIN=$(GOBIN) $(GO) install ./...
	pkg completion zsh > $(ZSH_COMPLETIONS)/_pkg
	utils completion zsh > $(ZSH_COMPLETIONS)/_utils

docs: build
	$(BIN)/pkg-$(GOOS)-$(GOARCH) docs --prefix docs/pkg
	$(BIN)/utils-$(GOOS)-$(GOARCH) docs --prefix docs/utils

update:
	$(GO) get -u ./...

fmt:
	$(GO) fmt ./...
	$(GO) mod tidy

$(BIN):
	mkdir -p $@
