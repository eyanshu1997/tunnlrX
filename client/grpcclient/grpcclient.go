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
	client proto.ConfigServiceClient
	Name   string
}

func NewGrpcClient(host string, port int, name string) (*grpcClient, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		err := fmt.Errorf("failed to connect to gRPC server: %v", err)
		serviceutils.Log.Error("Error: %s", err)
		return nil, err
	}
	defer conn.Close()

	serviceutils.Log.Info("Connected to gRPC server at %s:%d", host, port)
	client := proto.NewConfigServiceClient(conn)
	// Initialize and return a new gRPC client
	return &grpcClient{client: client, Name: name}, nil
}

func (g *grpcClient) RegisterClient(ctx context.Context) error {
	g.client.RegisterClient(ctx, &proto.RegisterClientRequest{
		Name: "TestClient",
	})
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
