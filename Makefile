# Project-level Makefile for TunnlrX
BUF_VERSION:=v1.17.0
SWAGGER_UI_VERSION:=v4.15.5


deps:
	@which go > /dev/null || (echo "Error: Go (golang) is not installed or not in PATH." && exit 1)
	@which protoc > /dev/null || (echo "Error: protoc is not installed or not in PATH." && exit 1)
	@which protoc-gen-go > /dev/null || go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@which protoc-gen-go-grpc > /dev/null || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest



all:  generate server client

generate: tidy proto swagger-ui

proto: deps
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate

swagger-ui:
	SWAGGER_UI_VERSION=$(SWAGGER_UI_VERSION) ./scripts/generate-swagger-ui.sh

tidy:
	go mod tidy

server: 
	cd server &&  go build -o ../build/tunnlrx-server

client: 
	cd client && go build -o ../build/tunnlrx-client

swagger:
	cd swagger-server && go build -o ../build/tunnelrx-swagger-server

clean:
	rm -f common/proto/*.pb.go
	rm -f build/tunnlrx-server build/tunnlrx-client
	cd server && go clean
	cd client && go clean


# we need to run as daemon
run_server:
	pkill tunnlrx-server || true
	nohup ./build/tunnlrx-server -config=configs/tunnlrx-server.json > logs/tunnlrx-server.log 2>&1 &

CLIENT_NO ?= 1

run_client:
	nohup ./build/tunnlrx-client -config=configs/tunnlrx-client.json > logs/tunnlrx-client$(CLIENT_NO).log 2>&1 &

run_swagger:
	nohup ./build/tunnelrx-swagger-server > logs/swagger-server.log 2>&1 &

stop_all:
	pkill tunnlrx-server || true
	pkill tunnlrx-client || true
	pkill tunnelrx-swagger-server || true