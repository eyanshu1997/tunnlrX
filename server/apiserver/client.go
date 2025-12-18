package apiserver

import (
	"context"
	"fmt"

	"github.com/eyanshu1997/tunnlrX/common/log"
	"github.com/eyanshu1997/tunnlrX/common/proto"
	"github.com/eyanshu1997/tunnlrX/server/mgmt"
	"github.com/eyanshu1997/tunnlrX/server/mgmttranslate"
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

func (s *TunnlrXApiServer) ListClients(ctx context.Context, req *proto.ListClientsRequest) (*proto.ListClientsResponse, error) {

	log.Info("ListClients called with request: %v", req)
	mgmtClients, err := mgmt.ListClients()
	if err != nil {
		log.Error("Error listing clients: %s", err)
		return nil, err
	}
	var clients []*proto.Client
	for _, client := range mgmtClients {
		clients = append(clients, mgmttranslate.CreateClientMsg(client))
	}
	return &proto.ListClientsResponse{
		Clients: clients,
	}, nil
}
