package main

import (
	"flag"

	"github.com/eyanshu1997/tunnlrX/common/log"
	"github.com/eyanshu1997/tunnlrX/server/apiserver"
	"github.com/eyanshu1997/tunnlrX/server/config"
	"github.com/eyanshu1997/tunnlrX/server/grpcserver"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "configs/tunnlrx-server.json", "Path to configuration file")
	flag.Parse()
}

func InitServer(config *config.ServerConfig) {

	go func() {
		log.Info("Starting GRPC server on port %d", config.GrpcPort)
		grpcServer, lis, err := grpcserver.GetGrpcServerAndListener(uint32(config.GrpcPort))
		if err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
	go func() {
		log.Info("Starting Api server on port %d", config.ApiPort)
		apiserver, lis, err := apiserver.GetApiServerandListener(uint32(config.ApiPort))
		if err != nil {
			log.Fatalf("Failed to start API server: %v", err)
		}
		if err := apiserver.Serve(lis); err != nil {
			log.Fatalf("Failed to serve API server: %v", err)
		}
	}()
}

func main() {
	if configPath == "" {
		panic("Config path is required")
	}

	config, err := config.LoadConfig(configPath)
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	log.InitLogger(config.LogLevel)

	InitServer(config)
	log.Info("Server started successfully")
	// listen for interrupt signal to gracefully shutdown the server
	select {}
}
