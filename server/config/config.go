package config

// This file contains configuration related code for the server
// we will be using json to configure becuase its humanreadable and easy to represent different types of data and easily parasble
// we will use a struct to represent the config and then use json package to parse it

import (
	"encoding/json"
	"os"

	"github.com/eyanshu1997/tunnlrx/common/serviceutils"
)

type ServerConfig struct {
	ApiPort  int `json:"port"`
	GrpcPort int `json:"grpc_port"`
	*serviceutils.ServiceConfig
}

func LoadConfig(filePath string) (*ServerConfig, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &ServerConfig{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
