# TunnlrX Server

This is the control plane for TunnlrX. It manages tunnels, domains, certificates, and exposes gRPC and HTTP APIs.

## Features
- gRPC API for tunnel management
- REST API via grpc-gateway

## Getting Started
1. Generate protobuf code
2. Run the server

## Dependencies
- Go 1.21+
- protoc
- google.golang.org/grpc
- github.com/grpc-ecosystem/grpc-gateway/v2
- google.golang.org/protobuf
