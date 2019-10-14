all: format test lint build

install:
	go build -o bin/go-guard

format:
	gofmt -w .

test:
	go test -cover ./...

lint:
	$(shell go env GOPATH)/bin/golangci-lint run --enable-all

godog:
	$(shell go env GOPATH)/bin/godog ./cmd/integration-tests

build:
	go build -o=bin/go-guard ./cmd/rest-server/