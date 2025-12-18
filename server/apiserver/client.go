package apiserver

import (
	"context"
	"fmt"

	"github.com/eyanshu1997/tunnlrX/common/log"
	"github.com/eyanshu1997/tunnlrX/common/proto"
	"github.com/eyanshu1997/tunnlrX/server/mgmt"
)

func (s *TunnlrXApiServer) CreateClient(ctx context.Context, req *proto.CreateClientRequest) (*proto.CreateClientResponse, error) {
	if req.GetName() == "" {
		return nil, fmt.Errorf("client name cannot be empty")
	}
	client, err := mgmt.NewClient(req.GetName())
	if err != nil {
		return nil, err
	}
	return &proto.CreateClientResponse{
		Id:        client.Id,
		Name:      client.Name,
		SecretKey: client.Secret,
	}, nil
}
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
