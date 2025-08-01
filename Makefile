GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/frochyzhang/ag-core/cmd/aif-go@latest
	go install github.com/frochyzhang/ag-core/cmd/protoc-gen-go-http@latest
	go install github.com/frochyzhang/ag-core/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	protoc-gen-go --version
	protoc-gen-go-grpc -version
	protoc-gen-go-http -version
	aif-go --version
	go get -u all

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

.PHONY: build build-macos build-linux build-windows
# Default build (current platform)
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

# Build for MacOS (darwin/arm64)
build-macos:
	GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/darwin-arm64/ ./...

# Build for Linux (linux/amd64)
build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/linux-amd64/ ./...

# Build for Windows (windows/amd64)
build-windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/windows-amd64/ ./...

# Build all platforms
build-all: build-macos build-linux build-windows

.PHONY: build-docker build-docker-linux
# build docker image for current platform
build-docker:
	export https_proxy=http://127.0.0.1:7890 http_proxy=http://127.0.0.1:7890 all_proxy=socks5://127.0.0.1:7890
	docker build -t ag-layout:$(VERSION) -f ./Dockerfile .
# build docker image for linux platform
build-docker-linux:
	export https_proxy=http://127.0.0.1:7890 http_proxy=http://127.0.0.1:7890 all_proxy=socks5://127.0.0.1:7890
	docker build -t ag-layout:$(VERSION) -f ./Dockerfile.linux .

.PHONY: generate
# generate
generate:
	go generate ./...
	go mod tidy

.PHONY: service
# service
service:
	aif-go proto service $(API_PROTO_FILES)

.PHONY: server
# server
server:
	aif-go proto server $(API_PROTO_FILES)

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;
# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
