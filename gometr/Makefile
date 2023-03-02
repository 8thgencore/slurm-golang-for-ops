install:
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin
	@GO111MODULE=on go mod download

run:
	@GO111MODULE=on go run cmd/server/main.go

lint:
	@GO111MODULE=on golangci-lint run ./... -v