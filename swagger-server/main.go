package main

import (
	"context"
	"flag"
	"fmt"
	"io/fs"
	"mime"
	"net/http"
	"strings"

	"github.com/eyanshu1997/tunnlrX/common/log"
	"github.com/eyanshu1997/tunnlrX/common/proto"
	"github.com/eyanshu1997/tunnlrX/swagger-server/config"
	"github.com/eyanshu1997/tunnlrX/third_party"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "configs/tunnlrx-client.json", "Path to configuration file")
	flag.Parse()
}

// getOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func getOpenAPIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}

// Run runs the gRPC-Gateway, dialling the provided address.
func Run(dialAddr string, port int) error {

	conn, err := grpc.Dial(
		dialAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		err := fmt.Errorf("failed to connect to gRPC server: %v", err)
		log.Error("Error: %s", err)
		return err
	}

	gwmux := runtime.NewServeMux()
	err = proto.RegisterTunnlrxApiServeiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}

	oa := getOpenAPIHandler()

	gatewayAddr := "0.0.0.0:" + fmt.Sprintf("%d", port)
	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				gwmux.ServeHTTP(w, r)
				return
			}
			oa.ServeHTTP(w, r)
		}),
	}

	log.Info("Serving gRPC-Gateway and OpenAPI Documentation on http://%s", gatewayAddr)
	return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServe())

}

func main() {
	swaggerServerConfig, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Failed to load client config: %v\n", err)
		return
	}
	log.InitLogger(swaggerServerConfig.LogLevel)
	if err := Run("dns:///"+swaggerServerConfig.ServerHost+":"+fmt.Sprintf("%d", swaggerServerConfig.ServerPort), swaggerServerConfig.UiPort); err != nil {
		grpclog.Fatal(err)
	}
}
