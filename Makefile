.PHONY: all proto

all: gateway
	./bin/buf generate --path ./proto/echoer.proto

gateway:
	go install github.com/bufbuild/buf/cmd/buf
