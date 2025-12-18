package grpcserver

import (
	"context"
	"fmt"
	"net"

	"github.com/eyanshu1997/tunnlrX/common/log"
	"github.com/eyanshu1997/tunnlrX/common/proto"
	"github.com/eyanshu1997/tunnlrX/server/mgmt"
	"google.golang.org/grpc/peer"
)

func (s *TunnlrxConfigServer) RegisterClient(ctx context.Context, req *proto.RegisterClientRequest) (*proto.RegisterClientResponse, error) {
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
	client, err := mgmt.GetClientByName(req.GetName())
	if err != nil {
		log.Error("Error getting client: %s", err)
		return nil, err
	}
	err = client.Register(clientIP, clientPort, req.GetSecretKey())
	if err != nil {
		log.Error("Error creating new client: %s", err)
		return nil, err
	}
	log.Info("Registered new client: %v", client)
	return &proto.RegisterClientResponse{
		Id: client.Id,
	}, nil
}
