package grpcserver

import (
	"context"
	"fmt"
	"net"

	"github.com/eyanshu1997/tunnlrx/common/log"
	"github.com/eyanshu1997/tunnlrx/common/proto"
	"google.golang.org/grpc/peer"
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
	log.Info("RegisterClient called with request: %v", req)
	var clientIP string
	var clientPort int
	if p, ok := peer.FromContext(ctx); ok {
		// To get the IP and port separately
		tcpAddr, ok := p.Addr.(*net.TCPAddr)
		if !ok {
			err := fmt.Errorf("unable to get the client ip and port")
			log.Error("Error: %s", err)
			return nil, err
		}
		clientIP = tcpAddr.IP.String()
		clientPort = tcpAddr.Port
		// Use clientIP and clientPort as needed
		log.Info("Client connected from IP: %s, Port: %d\n", clientIP, clientPort)

	}
	// check if existing client exists using same port and ip
	for _, client := range s.ClientDetails {
		if client.Ip == clientIP && client.Port == clientPort {
			log.Info("Client already exists: %v", client)
			return &proto.RegisterClientResponse{
				Id: client.Id,
			}, nil
		}
	}
	// create new client
	newClient := ClientDetails{
		Id:    tempid,
		Name:  req.GetName(),
		State: ClientStateActive,
		Ip:    clientIP,
		Port:  clientPort,
	}

	s.ClientDetails[newClient.Id] = newClient
	tempid++
	log.Info("Registered new client: %v", newClient)
	return &proto.RegisterClientResponse{
		Id: newClient.Id,
	}, nil
}

func (s *TunnelXServer) ListClients(ctx context.Context, req *proto.ListClientsRequest) (*proto.ListClientsResponse, error) {
	log.Info("ListClients called with request: %v", req)
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
