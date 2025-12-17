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

type ClientConfig struct {
	ServerHost   string         `json:"host"`
	ServerPort   int            `json:"port"`
	Name         string         `json:"name"`
	TunnelConfig []TunnelConfig `json:"tunnels"`
	LogLevel     string         `json:"log_level"`
}

func (s *ClientConfig) Validate() error {
	if s.ServerPort <= 0 || s.ServerPort > 65535 {
		return fmt.Errorf("error:Invalid Server port")
	}
	if s.ServerHost == "" {
		return fmt.Errorf("error:ServerHost cannot be empty")
	}
	if s.Name == "" {
		return fmt.Errorf("error:Name cannot be empty")
	}
	if len(s.TunnelConfig) == 0 {
		return fmt.Errorf("error:No tunnels configured")
	}
	if s.LogLevel == "" {
		return fmt.Errorf("invalid log level")
	}

	return nil
}

func (s *ClientConfig) String() string {
	return fmt.Sprintf("ClientConfig{ServerHost: %s, ServerPort: %d, Name: %s, TunnelConfig: %v, LogLevel: %s}", s.ServerHost, s.ServerPort, s.Name, s.TunnelConfig, s.LogLevel)
}

func LoadConfig(filePath string) (*ClientConfig, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &ClientConfig{}
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
