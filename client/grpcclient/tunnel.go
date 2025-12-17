package grpcclient

import (
	"context"
	"fmt"

	"github.com/eyanshu1997/tunnlrx/common/proto"
	"github.com/eyanshu1997/tunnlrx/common/serviceutils"
)

func (g *grpcClient) RegisterTunnel(ctx context.Context, port int32, domain string) error {
	_, err := g.client.RegisterTunnel(ctx, &proto.RegisterTunnelRequest{
		Name:     g.Name,
		Port:     port,
		ClientId: g.id,
		Domain:   domain,
	})
	if err != nil {
		return fmt.Errorf("failed to register client: %v", err)
	}
	return nil
}

func (g *grpcClient) ListTunnels(ctx context.Context) error {
	resp, err := g.client.ListTunnels(ctx, &proto.ListTunnelsRequest{})
	if err != nil {
		return fmt.Errorf("failed to list clients: %v", err)
	}
	serviceutils.Log.Info("ListTunnels response: %v", resp)
	return nil

}
