package grpcclient

import (
	"context"
	"fmt"

	"github.com/eyanshu1997/tunnlrx/common/proto"
	"github.com/eyanshu1997/tunnlrx/common/serviceutils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	// gRPC client fields
	conn   *grpc.ClientConn
	client proto.ConfigServiceClient
	Name   string
}

func NewGrpcClient(host string, port int, name string) (*grpcClient, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		err := fmt.Errorf("failed to connect to gRPC server: %v", err)
		serviceutils.Log.Error("Error: %s", err)
		return nil, err
	}

	serviceutils.Log.Info("Connected to gRPC server at %s:%d", host, port)
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

func (g *grpcClient) RegisterClient(ctx context.Context) error {
	_, err := g.client.RegisterClient(ctx, &proto.RegisterClientRequest{
		Name: g.Name,
	})
	if err != nil {
		return fmt.Errorf("failed to register client: %v", err)
	}
	return nil
}

func (g *grpcClient) ListClients(ctx context.Context) error {
	resp, err := g.client.ListClients(ctx, &proto.ListClientsRequest{})
	if err != nil {
		return fmt.Errorf("failed to list clients: %v", err)
	}
	serviceutils.Log.Info("ListClients response: %v", resp)
	return nil
}
