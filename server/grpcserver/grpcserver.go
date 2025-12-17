package grpcserver

import (
	"fmt"
	"net"

	"github.com/eyanshu1997/tunnlrx/common/proto"
	"github.com/eyanshu1997/tunnlrx/common/serviceutils"
	"google.golang.org/grpc"
)

type TunnelXServer struct {
	// Add fields as necessary, e.g., database connections, configurations, etc.
	ClientDetails map[uint32]ClientDetails // Example field to hold tunnel states
	TunnelDetails map[uint32]TunnelDetails // Example field to hold tunnel states
	proto.UnimplementedConfigServiceServer
}

func NewTunnelXServer() *TunnelXServer {
	return &TunnelXServer{
		ClientDetails: make(map[uint32]ClientDetails),
		TunnelDetails: make(map[uint32]TunnelDetails),
	}
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
