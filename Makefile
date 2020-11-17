.PHONY: all proto

export GOBIN := $(CURDIR)/bin
export PATH := $(GOBIN):$(PATH)

PROTO_SOURCES := $(shell find ./ -maxdepth 1 -type f -name '*.proto')
PROTO_TARGETS := $(addsuffix .pb.go,$(basename $(PROTO_SOURCES)))

all: proto

proto: $(PROTO_TARGETS)

%.pb.go: gateway %.proto
	protoc \
		-I=$(CURDIR) \
		-I=$(GOPATH)/pkg/mod \
		--go_out . --go_opt paths=source_relative \
		--go-grpc_out . --go-grpc_opt paths=source_relative \
		--grpc-gateway_out . \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		--openapiv2_out . --openapiv2_opt logtostderr=true \
		$(CURDIR)/$(*).proto

gateway: go.sum
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc

go.sum: go.mod
	go mod tidy