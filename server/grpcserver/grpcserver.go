package grpcserver

import (
	"context"
	"fmt"
	"net"

	"github.com/eyanshu1997/tunnlrx/common/proto"
	"github.com/eyanshu1997/tunnlrx/common/serviceutils"
	"google.golang.org/grpc"
)

type TunnelDetail struct {
	ID     int
	Name   string
	Status string
	// Add other fields as necessary
}

type TunnelXServer struct {
	// Add fields as necessary, e.g., database connections, configurations, etc.
	TunnelDetails map[int]TunnelDetail // Example field to hold tunnel states
	proto.UnimplementedTunnelServiceServer
}

func NewTunnelXServer() *TunnelXServer {
	return &TunnelXServer{
		TunnelDetails: make(map[int]TunnelDetail),
	}
}

// Implement the TunnelServiceServer interface methods here.
func (s *TunnelXServer) CreateTunnel(ctx context.Context, req *proto.CreateTunnelRequest) (*proto.CreateTunnelResponse, error) {
	id := len(s.TunnelDetails) + 1
	if s.TunnelDetails == nil {
		s.TunnelDetails = make(map[int]TunnelDetail)
	}
	s.TunnelDetails[id] = TunnelDetail{
		ID:     id,
		Name:   req.Name,
		Status: "active",
	}
	return &proto.CreateTunnelResponse{
		Id: fmt.Sprintf("%d", id),
	}, nil
}

func (s *TunnelXServer) ListTunnels(ctx context.Context, req *proto.ListTunnelsRequest) (*proto.ListTunnelsResponse, error) {
	var tunnels []*proto.Tunnel
	for _, detail := range s.TunnelDetails {
		tunnels = append(tunnels, &proto.Tunnel{
			Id:   fmt.Sprintf("%d", detail.ID),
			Name: detail.Name,
		})
	}
	return &proto.ListTunnelsResponse{
		Tunnels: tunnels,
	}, nil
}

func GetGrpcServerAndListener(port uint32) (*grpc.Server, net.Listener, error) {
	var opts []grpc.ServerOption

	// initialize gRPC server
	grpcServer := grpc.NewServer(opts...)
	tunnelXServer := NewTunnelXServer()
	serviceutils.Log.Debug("Registering TunnelServiceServer %s with gRPC server %s ", tunnelXServer, grpcServer)
	proto.RegisterTunnelServiceServer(grpcServer, tunnelXServer)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		serviceutils.Log.Fatalf("Failed to listen on port %d: %v", port, err)
		return nil, nil, err
	}
	return grpcServer, lis, nil
}
