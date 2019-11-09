all: format lint godog

format:
	gofmt -w .

lint:
	$(shell go env GOPATH)/bin/golangci-lint run --enable-all

godog:
	$(shell go env GOPATH)/bin/godog ./features
