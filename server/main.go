package main

import (
	"net/http"

	"flag"

	_ "github.com/eyanshu1997/tunnlrx/common/proto"
	"github.com/eyanshu1997/tunnlrx/common/serviceutils"
	"github.com/eyanshu1997/tunnlrx/server/config"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "config.json", "Path to configuration file")
	flag.Parse()
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
	serviceutils.Log.Info("🚀 TunnlrX server starting on :8080")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	http.ListenAndServe(":8080", nil)
}
