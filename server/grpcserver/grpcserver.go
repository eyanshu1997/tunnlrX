package grpcserver

import (
	"fmt"
	"net"

	"github.com/eyanshu1997/tunnlrX/common/log"
	"github.com/eyanshu1997/tunnlrX/common/proto"
	"google.golang.org/grpc"
)

type TunnlrxConfigServer struct {
	proto.UnimplementedConfigServiceServer
}

func NewTunnlrxConfigServer() *TunnlrxConfigServer {
	return &TunnlrxConfigServer{}
}

func GetGrpcServerAndListener(port uint32) (*grpc.Server, net.Listener, error) {
	var opts []grpc.ServerOption

	// initialize gRPC server
	grpcServer := grpc.NewServer(opts...)
	TunnlrxConfigServer := NewTunnlrxConfigServer()
	log.Debug("Registering TunnelServiceServer %s with gRPC server %s ", TunnlrxConfigServer, grpcServer)
	proto.RegisterConfigServiceServer(grpcServer, TunnlrxConfigServer)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
		return nil, nil, err
	}
	return grpcServer, lis, nil
}
