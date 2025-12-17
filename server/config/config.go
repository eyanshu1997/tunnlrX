package config

// This file contains configuration related code for the server
// we will be using json to configure because its human readable and easy to represent different types of data and easily parasble
// we will use a struct to represent the config and then use json package to parse it

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServerConfig struct {
	ApiPort  int    `json:"port"`
	GrpcPort int    `json:"grpc_port"`
	LogLevel string `json:"log_level"`
}

func (s *ServerConfig) Validate() error {
	if s.ApiPort <= 0 || s.ApiPort > 65535 {
		return fmt.Errorf("invalid API port")
	}
	if s.GrpcPort <= 0 || s.GrpcPort > 65535 {
		return fmt.Errorf("invalid gRPC port")
	}
	if s.LogLevel == "" {
		return fmt.Errorf("invalid log level")
	}
	return nil
}

func (s *ServerConfig) String() string {
	return fmt.Sprintf("ServerConfig{ApiPort: %d, GrpcPort: %d, LogLevel: %s}", s.ApiPort, s.GrpcPort, s.LogLevel)
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
	err = config.Validate()
	if err != nil {
		return nil, err
	}

	return config, nil
}
