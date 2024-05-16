MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
.DELETE_ON_ERROR:
.SUFFIXES:

APP ?= glci-linter
PROJECT ?= gitlab.com/eyecon/vavada/devops/$(APP)

GIT_COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%F_%T')

# Go build flags
GOOS ?= linux
GOARCH ?= amd64
GOLDFLAGS := '-w -s -extldflags "-static" -X ${PROJECT}/version.Release="${GIT_COMMIT}" -X ${PROJECT}/version.Commit="${GIT_COMMIT}" -X ${PROJECT}/version.BuildTime="${BUILD_TIME}"'

#help:
#	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Help target
help:
	@echo ''
	@echo 'Usage: make [TARGET]'
	@echo 'Targets:'
	@echo '  help     	display this message'
	@echo '  fmt      	gofmt vendor'
	@echo '  test     	run go test'
	@echo '  lint     	run go linter'
	@echo '  mod           run go mod download'
	@echo '  build        run go build'
	@echo '  all     	run go fmt lint (default make)'
	@echo ''

.PHONY: help

# Dynamic targets
PLATFORMS ?= linux/amd64 linux/arm64 darwin/amd64 darwin/arm64

temp = $(subst /, ,$@)
GOOS = $(word 1, $(temp))
GOARCH = $(word 2, $(temp))

.PHONY: all
all: fmt lint test

.PHONY: $(PLATFORMS) build
build: $(PLATFORMS)
$(PLATFORMS):
	@echo "-> $@"
	@echo "Building Go binary"
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -trimpath -ldflags $(GOLDFLAGS) -installsuffix cgo -o build/glci-linter_$(GOOS)_$(GOARCH) .

.PHONY: test
test:
	@echo "-> $@"
	go test -v -race ./...

.PHONY: fmt
fmt:
	@echo "-> $@"
	@gofmt -s -l ./ | grep -v vendor | tee /dev/stderr

.PHONY: lint
lint:
	@echo "-> $@"
	@go get -u golang.org/x/lint/golint
	@golint ./... | tee /dev/stderr
	@go get -u golang.org/x/tools/go/analysis/cmd/vet
	@go vet --all

.PHONY: mod
mod:
	@echo "-> $@"
	@go mod download

.PHONY: clean
clean:
	@echo "-> $@"
	@rm -Rf ./build
