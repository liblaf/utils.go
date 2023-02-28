BIN             := $(CURDIR)/bin
SHELL           := /bin/bash
ZSH_COMPLETIONS := $(ZSH_CUSTOM)/plugins/completions

GOARCH ?= $(shell go env GOARCH)
GOBIN  ?= $(HOME)/.local/bin
GOEXE  ?= $(shell go env GOEXE)
GOOS   ?= $(shell go env GOOS)

ENV := env GOARCH=$(GOARCH) GOBIN=$(GOBIN) GOOS=$(GOOS)
GO  := $(ENV) go

build: | $(BIN)
	$(GO) build -o $(BIN) ./...

clean:
	$(GO) clean
	$(RM) -r $(BIN)
	$(RM) docs/hello/hello*.md
	$(RM) docs/pkg/pkg*.md
	$(RM) docs/utils/utils*.md

docs: build
	$(BIN)/hello$(GOEXE) docs -p docs/hello
	$(BIN)/pkg$(GOEXE) docs -p docs/pkg
	$(BIN)/utils$(GOEXE) docs -p docs/utils

fmt:
	$(GO) fmt ./...
	$(GO) mod tidy

install:
	$(GO) install ./...
	pkg completion zsh > $(ZSH_COMPLETIONS)/_pkg
	utils completion zsh > $(ZSH_COMPLETIONS)/_utils

rename: build
	$(ENV) $(SHELL) $(CURDIR)/scripts/rename.sh

update:
	$(GO) get -u ./...

$(BIN):
	mkdir -p $@
