package grpcserver

import (
	"context"

	"github.com/eyanshu1997/tunnlrx/common/proto"
	"github.com/eyanshu1997/tunnlrx/common/serviceutils"
)

type ClientState uint32

const (
	ClientStateUnknown ClientState = iota
	ClientStateInactive
	ClientStateActive
	ClientStateInprogress
)

type ClientDetails struct {
	Id    uint32
	Name  string
	Port  int
	Ip    string
	State ClientState
}

var tempid uint32 = 0

func (s *TunnelXServer) RegisterClient(ctx context.Context, req *proto.RegisterClientRequest) (*proto.RegisterClientResponse, error) {
	serviceutils.Log.Info("RegisterClient called with request: %v", req)
	// TODO get ip and port details from the context metadata

	newClient := ClientDetails{
		Id:    tempid,
		Name:  req.GetName(),
		State: ClientStateActive,
	}

	s.ClientDetails[newClient.Id] = newClient
	tempid++
	serviceutils.Log.Info("Registered new client: %v", newClient)
	return &proto.RegisterClientResponse{
		Id: newClient.Id,
	}, nil
}

func (s *TunnelXServer) ListClients(ctx context.Context, req *proto.ListClientsRequest) (*proto.ListClientsResponse, error) {
	serviceutils.Log.Info("ListClients called with request: %v", req)
	var clients []*proto.Client
	for _, client := range s.ClientDetails {
		clients = append(clients, &proto.Client{
			Id:          client.Id,
			Name:        client.Name,
			ClientState: proto.ClientState(client.State),
		})
	}
	return &proto.ListClientsResponse{
		Clients: clients,
	}, nil
}
