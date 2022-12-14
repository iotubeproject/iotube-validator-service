GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

build: build-validator build-relayer build-query

build-query:
	$(GOBUILD) -o ./bin/query cmd/query/main.go

build-relayer:
	$(GOBUILD) -o ./bin/relayer cmd/relayer/main.go

build-validator:
	$(GOBUILD) -o ./bin/validator cmd/validator/main.go

proto:
	protoc -I . -I ./api --go_out=. --go_opt=paths=source_relative --grpc-gateway_out paths=source_relative:. --go-grpc_out=. --go-grpc_opt=paths=source_relative api/service.proto

.PHONY: build proto
