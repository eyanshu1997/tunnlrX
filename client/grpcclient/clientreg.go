package grpcclient

import (
	"context"
	"fmt"

	"github.com/eyanshu1997/tunnlrx/common/log"
	"github.com/eyanshu1997/tunnlrx/common/proto"
)

func (g *grpcClient) RegisterClient(ctx context.Context) error {
	resp, err := g.client.RegisterClient(ctx, &proto.RegisterClientRequest{
		Name: g.Name,
	})
	if err != nil {
		return fmt.Errorf("failed to register client: %v", err)
	}
	g.id = resp.GetId()
	log.Info("RegisterClient got id: %d", g.id)
	return nil
}

func (g *grpcClient) ListClients(ctx context.Context) error {
	resp, err := g.client.ListClients(ctx, &proto.ListClientsRequest{})
	if err != nil {
		return fmt.Errorf("failed to list clients: %v", err)
	}
	log.Info("ListClients response: %v", resp)
	return nil
}
