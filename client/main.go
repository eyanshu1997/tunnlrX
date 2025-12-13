package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/eyanshu1997/tunnlrx/client/config"
	"github.com/eyanshu1997/tunnlrx/client/grpcclient"
	"github.com/eyanshu1997/tunnlrx/common/serviceutils"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "config.json", "Path to configuration file")
	flag.Parse()
}

func main() {

	// Load client configuration
	clientConfig, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Failed to load client config: %v\n", err)
		return
	}
	serviceutils.InitServiceUtils(clientConfig.ServiceConfig, "TunnlrxServer")
	// Initialize client with the loaded configuration
	client, err := grpcclient.NewGrpcClient(clientConfig.ServerHost, clientConfig.ServerPort, clientConfig.Name)
	if err != nil {
		fmt.Printf("Failed to initialize client: %v\n", err)
		return
	}
	defer func() {
		_ = client.Close()
	}()
	serviceutils.Log.Info("Client initialized successfully:", client)
	ctx := serviceutils.GetContextWithMetadata()
	// client will register itself and list clients from server every 30 seconds here
	err = client.RegisterClient(ctx)
	if err != nil {
		fmt.Printf("Failed to register client: %v\n", err)
		return
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
