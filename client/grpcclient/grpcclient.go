package grpcclient

import (
	"fmt"

	"github.com/eyanshu1997/tunnlrx/common/log"
	"github.com/eyanshu1997/tunnlrx/common/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	// gRPC client fields
	conn   *grpc.ClientConn
	client proto.ConfigServiceClient
	Name   string
	id     uint32
}

func NewGrpcClient(host string, port int, name string) (*grpcClient, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		err := fmt.Errorf("failed to connect to gRPC server: %v", err)
		log.Error("Error: %s", err)
		return nil, err
	}

	log.Info("Connected to gRPC server at %s:%d", host, port)
	client := proto.NewConfigServiceClient(conn)
	// Initialize and return a new gRPC client (keep connection open until Close is called)
	return &grpcClient{conn: conn, client: client, Name: name}, nil
}

// Close closes the underlying gRPC connection.
func (g *grpcClient) Close() error {
	if g.conn != nil {
		return g.conn.Close()
	}
	return nil
}
