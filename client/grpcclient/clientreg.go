package grpcclient

import (
	"context"
	"fmt"

	"github.com/eyanshu1997/tunnlrX/common/log"
	"github.com/eyanshu1997/tunnlrX/common/proto"
)

func (g *grpcClient) RegisterClient(ctx context.Context) error {
	resp, err := g.client.RegisterClient(ctx, &proto.RegisterClientRequest{
		Name:      g.Name,
		SecretKey: g.Secret,
	})
	if err != nil {
		return fmt.Errorf("failed to register client: %v", err)
	}
	g.id = resp.GetId()
	log.Info("RegisterClient got id: %d", g.id)
	return nil
}
