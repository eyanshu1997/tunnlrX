package grpcserver

import (
	"context"
	"fmt"

	"github.com/eyanshu1997/tunnlrx/common/proto"
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
	*proto.UnimplementedTunnelServiceServer
}

func NewTunnelXServer() *TunnelXServer {
	return &TunnelXServer{}
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

func GetGrpcServer() *grpc.Server {
	// initialize gRPC server
	grpcServer := grpc.NewServer()
	tunnelXServer := NewTunnelXServer()
	proto.RegisterTunnelServiceServer(grpcServer, tunnelXServer)
	return grpcServer
}
