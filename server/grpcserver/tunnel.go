package grpcserver

import (
	"context"
	"fmt"

	"github.com/eyanshu1997/tunnlrX/common/log"
	"github.com/eyanshu1997/tunnlrX/common/proto"
	"github.com/eyanshu1997/tunnlrX/server/mgmt"
	"github.com/eyanshu1997/tunnlrX/server/mgmttranslate"
)

func (s *TunnlrxConfigServer) ListTunnels(ctx context.Context, req *proto.ListTunnelsRequest) (*proto.ListTunnelsResponse, error) {
	log.Info("ListTunnels called with request: %v", req)
	// check if optional client id is sent
	var tunnels []*proto.Tunnel
	mgmtTunnels, err := mgmt.ListTunnels()
	if err != nil {
		log.Error("Error listing tunnels: %s", err)
		return nil, err
	}
	for _, tunnel := range mgmtTunnels {
		mgmtClient, err := mgmt.GetClient(tunnel.ClientId)
		if err != nil {
			err := fmt.Errorf("client id %d not found for tunnel id %d", tunnel.ClientId, tunnel.Id)
			log.Error("Error: %s", err)
			return nil, err
		}
		tunnels = append(tunnels, mgmttranslate.CreateTunnelMsg(tunnel, mgmtClient))
	}
	return &proto.ListTunnelsResponse{
		Tunnels: tunnels,
	}, nil
}
