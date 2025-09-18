# Project-level Makefile for TunnlrX

PROTO_SRC := common/proto/tunnel.proto

PROTOC_GEN_GO := protoc --go_out=paths=source_relative:common/proto --go-grpc_out=paths=source_relative:common/proto

.PHONY: all proto server client clean

all: proto server client

proto:
	$(PROTOC_GEN_GO) $(PROTO_SRC)

server:
	cd server && go build -o ../build/tunnlrx-server

client:
	cd client && go build -o ../build/tunnlrx-client

clean:
	rm -f common/proto/*.pb.go
	rm -f build/tunnlrx-server build/tunnlrx-client
	cd server && go clean
	cd client && go clean
