package apiserver

import (
	"fmt"
	"net"

	"github.com/eyanshu1997/tunnlrX/common/log"
	"github.com/eyanshu1997/tunnlrX/common/proto"
	"google.golang.org/grpc"
)

type TunnlrXApiServer struct {
	proto.UnimplementedTunnlrxApiServeiceServer
}

func NewTunnlrXApiServer() *TunnlrXApiServer {
	return &TunnlrXApiServer{}
}

func GetApiServerandListener(port uint32) (*grpc.Server, net.Listener, error) {
	var opts []grpc.ServerOption

	// initialize gRPC server
	grpcServer := grpc.NewServer(opts...)
	TunnlrXApiServer := NewTunnlrXApiServer()
	log.Debug("Registering TunnlrXApiServeiceServer %s with gRPC server %s ", TunnlrXApiServer, grpcServer)
	proto.RegisterTunnlrxApiServeiceServer(grpcServer, TunnlrXApiServer)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
		return nil, nil, err
	}
	return grpcServer, lis, nil
}
