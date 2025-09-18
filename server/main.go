package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/eyanshu1997/tunnlrx/common/serviceutils"
	"github.com/eyanshu1997/tunnlrx/server/config"
	"github.com/eyanshu1997/tunnlrx/server/grpcserver"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "config.json", "Path to configuration file")
	flag.Parse()
}

func InitServer(config *config.ServerConfig) {
	httpMux := http.NewServeMux()
	grpcServer := grpcserver.GetGrpcServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcPort))
	if err != nil {
		serviceutils.Log.Fatalf("Failed to listen on port %d: %v", config.GrpcPort, err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			serviceutils.Log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// Simple HTTP server for health checks
	httpMux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	serviceutils.Log.Info("Starting HTTP server on port %d", config.ApiPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), httpMux); err != nil {
		serviceutils.Log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func main() {
	if configPath == "" {
		panic("Config path is required")
	}

	config, err := config.LoadConfig(configPath)
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	serviceutils.InitServiceUtils(config.ServiceConfig)
	serviceutils.Log.Info("🚀 TunnlrX server starting on ")
	InitServer(config)
}
