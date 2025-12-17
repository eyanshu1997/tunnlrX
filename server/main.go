package main

import (
	"flag"
	"net/http"

	"github.com/eyanshu1997/tunnlrx/common/log"
	"github.com/eyanshu1997/tunnlrx/server/config"
	"github.com/eyanshu1997/tunnlrx/server/grpcserver"
	"github.com/eyanshu1997/tunnlrx/server/httpserver"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "configs/tunnlrx-server.json", "Path to configuration file")
	flag.Parse()
}

func InitServer(config *config.ServerConfig) {

	go func() {
		log.Log.Info("Starting GRPC server on port %d", config.GrpcPort)
		grpcServer, lis, err := grpcserver.GetGrpcServerAndListener(uint32(config.GrpcPort))
		if err != nil {
			log.Log.Fatalf("Failed to start gRPC server: %v", err)
		}
		if err := grpcServer.Serve(lis); err != nil {
			log.Log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	go func() {
		log.Log.Info("Starting HTTP server on port %d", config.ApiPort)
		httpServer := httpserver.NewHttpServer(config.ApiPort)
		if err := httpServer.Start(); err != nil && err != http.ErrServerClosed {
			log.Log.Fatalf("Failed to start HTTP server: %v", err)
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
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	InitServer(config)
	log.Log.Info("Server started successfully")
	// listen for interrupt signal to gracefully shutdown the server
	select {}
}
