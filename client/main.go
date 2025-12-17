package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/eyanshu1997/tunnlrx/client/config"
	"github.com/eyanshu1997/tunnlrx/client/grpcclient"
	"github.com/eyanshu1997/tunnlrx/common/log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "configs/tunnlrx-client.json", "Path to configuration file")
	flag.Parse()
}

func main() {

	// Load client configuration
	clientConfig, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Failed to load client config: %v\n", err)
		return
	}

	log.InitLogger(clientConfig.LogLevel)

	// Initialize client with the loaded configuration
	client, err := grpcclient.NewGrpcClient(clientConfig.ServerHost, clientConfig.ServerPort, clientConfig.Name)
	if err != nil {
		fmt.Printf("Failed to initialize client: %v\n", err)
		return
	}
	defer func() {
		_ = client.Close()
	}()
	log.Info("Client initialized successfully:", client)

	// all of this code will be changed to actual implementation
	// client will register itself and list clients from server every 30 seconds here
	ctx := context.Background()
	err = client.RegisterClient(ctx)
	if err != nil {
		fmt.Printf("Failed to register client: %v\n", err)
		return
	}
	for _, tunnel := range clientConfig.TunnelConfig {
		err = client.RegisterTunnel(ctx, int32(tunnel.Port), tunnel.Domain)
		if err != nil {
			fmt.Printf("Failed to register tunnel: %v\n", err)
			return
		}
	}

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		err := client.ListClients(ctx)
		if err != nil {
			fmt.Printf("Failed to list clients: %v\n", err)
		}
	}

}
