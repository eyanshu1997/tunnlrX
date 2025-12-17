package grpcserver

import (
	"context"
	"fmt"

	"github.com/eyanshu1997/tunnlrx/common/log"
	"github.com/eyanshu1997/tunnlrx/common/proto"
)

type TunnelState uint32

const (
	TunnelStateUnknown TunnelState = iota
	TunnelStateInactive
	TunnelStateActive
	TunnelStateInprogress
)

type TunnelDetails struct {
	Id       uint32
	Name     string
	Port     int
	State    TunnelState
	ClientId uint32
}

func (s *TunnelXServer) RegisterTunnel(ctx context.Context, req *proto.RegisterTunnelRequest) (*proto.RegisterTunnelResponse, error) {
	log.Info("RegisterTunnel called with request: %v", req)
	newTunnel := TunnelDetails{
		Id:       tempid,
		Name:     req.GetName(),
		State:    TunnelStateActive,
		ClientId: req.GetClientId(),
	}
	// check if client id is present
	if _, ok := s.ClientDetails[newTunnel.ClientId]; !ok {
		err := fmt.Errorf("client id %d not found", newTunnel.ClientId)
		log.Error("Error: %s", err)
		return nil, err
	}
	s.TunnelDetails[newTunnel.Id] = newTunnel
	tempid++
	log.Info("Registered new tunnel: %v", newTunnel)
	return &proto.RegisterTunnelResponse{
		Id: newTunnel.Id,
	}, nil
}

func (s *TunnelXServer) ListTunnels(ctx context.Context, req *proto.ListTunnelsRequest) (*proto.ListTunnelsResponse, error) {
	log.Info("ListTunnels called with request: %v", req)
	// check if optional client id is sent
	var tunnels []*proto.Tunnel
	for _, tunnel := range s.TunnelDetails {
		if req.GetClientId() != 0 && tunnel.ClientId != req.GetClientId() {
			continue
		}
		client, ok := s.ClientDetails[tunnel.ClientId]
		if !ok {
			err := fmt.Errorf("client id %d not found", tunnel.ClientId)
			log.Error("Error: %s", err)
			return nil, err
		}
		tunnels = append(tunnels, &proto.Tunnel{
			Id:          tunnel.Id,
			Name:        tunnel.Name,
			TunnelState: proto.TunnelState(tunnel.State),
			Client: &proto.Client{
				Id:          client.Id,
				Name:        client.Name,
				ClientState: proto.ClientState(client.State),
			},
		})

	}
	return &proto.ListTunnelsResponse{
		Tunnels: tunnels,
	}, nil
}
