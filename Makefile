.PHONY: all proto

export GOBIN := $(CURDIR)/bin
export PATH := $(GOBIN):$(PATH)

all: gateway
	./bin/buf generate --path ./proto/echoer.proto

gateway: go.sum
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/bufbuild/buf/cmd/buf

go.sum: go.mod
	go mod tidy