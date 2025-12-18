package grpcclient

import (
	"context"
	"fmt"

	"github.com/eyanshu1997/tunnlrX/common/log"
	"github.com/eyanshu1997/tunnlrX/common/proto"
)

func (g *grpcClient) ListTunnels(ctx context.Context) error {
	resp, err := g.client.ListTunnels(ctx, &proto.ListTunnelsRequest{})
	if err != nil {
		return fmt.Errorf("failed to list clients: %v", err)
	}
	log.Info("ListTunnels response: %v", resp)
	return nil

}
