package grpcserver

import (
	"context"
	"fmt"
	"net"

	"github.com/eyanshu1997/tunnlrx/common/proto"
	"github.com/eyanshu1997/tunnlrx/common/serviceutils"
	"google.golang.org/grpc"
)

type ClientStatus int

const (
	Unknown ClientStatus = iota
	Inactive
	Active
	Inprogress
)

type ClientDetails struct {
	Id     uint32
	Name   string
	Port   int
	Ip     string
	Domain string
	Status ClientStatus
}

type TunnelXServer struct {
	// Add fields as necessary, e.g., database connections, configurations, etc.
	TunnelDetails map[uint32]ClientDetails // Example field to hold tunnel states
	proto.UnimplementedConfigServiceServer
}

func NewTunnelXServer() *TunnelXServer {
	return &TunnelXServer{
		TunnelDetails: make(map[uint32]ClientDetails),
	}
}

func GenerateDomainName(clientId uint32) string {
	return fmt.Sprintf("client-%d.tunnlrx.example.com", clientId)
}

var tempid uint32 = 0

func (s *TunnelXServer) RegisterClient(ctx context.Context, req *proto.RegisterClientRequest) (*proto.RegisterClientResponse, error) {
	serviceutils.Log.Info("RegisterClient called with request: %v", req)
	// TODO get ip and port details from the context metadata

	newClient := ClientDetails{
		Id:     tempid,
		Name:   req.GetName(),
		Domain: GenerateDomainName(tempid),
		Status: Inprogress,
	}

	s.TunnelDetails[newClient.Id] = newClient
	tempid++
	serviceutils.Log.Info("Registered new client: %v", newClient)
	return &proto.RegisterClientResponse{
		Id:     newClient.Id,
		Domain: newClient.Domain,
	}, nil
}

func (s *TunnelXServer) ListClients(ctx context.Context, req *proto.ListClientsRequest) (*proto.ListClientsResponse, error) {
	serviceutils.Log.Info("ListClients called with request: %v", req)
	// Implement the logic to list clients
	return &proto.ListClientsResponse{}, nil
}

func GetGrpcServerAndListener(port uint32) (*grpc.Server, net.Listener, error) {
	var opts []grpc.ServerOption

	// initialize gRPC server
	grpcServer := grpc.NewServer(opts...)
	tunnelXServer := NewTunnelXServer()
	serviceutils.Log.Debug("Registering TunnelServiceServer %s with gRPC server %s ", tunnelXServer, grpcServer)
	proto.RegisterConfigServiceServer(grpcServer, tunnelXServer)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		serviceutils.Log.Fatalf("Failed to listen on port %d: %v", port, err)
		return nil, nil, err
	}
	return grpcServer, lis, nil
}
