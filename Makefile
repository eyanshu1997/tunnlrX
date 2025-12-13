# Project-level Makefile for TunnlrX

PROTO_SRC := common/proto/config.proto

PROTOC_GEN_GO := protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:.

.PHONY: all proto server client clean deps

all:  proto server client


deps:
	@which go > /dev/null || (echo "Error: Go (golang) is not installed or not in PATH." && exit 1)
	@which protoc > /dev/null || (echo "Error: protoc is not installed or not in PATH." && exit 1)
	@which protoc-gen-go > /dev/null || go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@which protoc-gen-go-grpc > /dev/null || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

proto: deps
	cd common && go mod tidy
	$(PROTOC_GEN_GO) $(PROTO_SRC)

server: 
	cd server && go mod tidy && go build -o ../build/tunnlrx-server

client: 
	cd client && go mod tidy && go build -o ../build/tunnlrx-client

qserver: 
	cd server && go build -o ../build/tunnlrx-server

qclient: 
	cd client  && go build -o ../build/tunnlrx-client

clean:
	rm -f common/proto/*.pb.go
	rm -f build/tunnlrx-server build/tunnlrx-client
	cd server && go clean
	cd client && go clean


# we need to run as daemon
run_server:
	pkill tunnlrx-server || true
	nohup ./build/tunnlrx-server -config=configs/tunnlrx-server.json > tunnlrx-server.log 2>&1 &

run_client:
	pkill tunnlrx-client || true
	nohup ./build/tunnlrx-client -config=configs/tunnlrx-client.json > tunnlrx-client.log 2>&1 &

stop_all:
	pkill tunnlrx-server || true
	pkill tunnlrx-client || true