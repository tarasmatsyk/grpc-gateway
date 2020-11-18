.PHONY: all proto

all: gateway
	./bin/buf generate --path ./proto/echoer

gateway:
	go install github.com/bufbuild/buf/cmd/buf
