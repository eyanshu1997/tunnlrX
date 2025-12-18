package config

// This file contains configuration related code for the server
// we will be using json to configure because its human readable and easy to represent different types of data and easily parasble
// we will use a struct to represent the config and then use json package to parse it

import (
	"encoding/json"
	"fmt"
	"os"
)

type TunnelConfig struct {
	Port   uint32 `json:"port"`
	Domain string `json:"domain"`
}

type SwaggerServerConfig struct {
	ServerHost string `json:"host"`
	ServerPort int    `json:"port"`
	LogLevel   string `json:"log_level"`
	UiPort     int    `json:"ui_port"`
}

func (s *SwaggerServerConfig) Validate() error {
	if s.ServerPort <= 0 || s.ServerPort > 65535 {
		return fmt.Errorf("error:Invalid Server port")
	}
	if s.ServerHost == "" {
		return fmt.Errorf("error:ServerHost cannot be empty")
	}

	if s.LogLevel == "" {
		return fmt.Errorf("invalid log level")
	}

	return nil
}

func (s *SwaggerServerConfig) String() string {
	return fmt.Sprintf("ServerHost: %s, ServerPort: %d, LogLevel: %s", s.ServerHost, s.ServerPort, s.LogLevel)
}

func LoadConfig(filePath string) (*SwaggerServerConfig, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &SwaggerServerConfig{}
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
