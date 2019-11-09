all: format lint test

format:
	gofmt -w .

lint:
	$(shell go env GOPATH)/bin/golangci-lint run --enable-all

test:
    go test -v -cover ./...

godog:
	$(shell go env GOPATH)/bin/godog ./features
