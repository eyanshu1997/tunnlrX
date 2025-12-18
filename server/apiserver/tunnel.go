package apiserver

import (
	"context"

	"github.com/eyanshu1997/tunnlrX/common/log"
	"github.com/eyanshu1997/tunnlrX/common/proto"
	"github.com/eyanshu1997/tunnlrX/server/mgmt"
)

func (s *TunnlrXApiServer) CreateTunnel(ctx context.Context, req *proto.CreateTunnelRequest) (*proto.CreateTunnelResponse, error) {

	log.Info("CreateTunnel called with request: %v", req)
	newTunnel := mgmt.TunnelDetails{
		Name:     req.GetName(),
		ClientId: req.GetClientId(),
	}
	err := newTunnel.Create()
	if err != nil {
		log.Error("Error creating new tunnel: %s", err)
		return nil, err
	}
	log.Info("Registered new tunnel: %v", newTunnel)
	return &proto.CreateTunnelResponse{
		Id: newTunnel.Id,
	}, nil
}
