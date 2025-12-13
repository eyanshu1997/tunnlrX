package config

// This file contains configuration related code for the server
// we will be using json to configure becuase its humanreadable and easy to represent different types of data and easily parasble
// we will use a struct to represent the config and then use json package to parse it

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/eyanshu1997/tunnlrx/common/serviceutils"
)

type ClientConfig struct {
	ServerHost string `json:"host"`
	ServerPort int    `json:"port"`
	Name       string `json:"name"`
	*serviceutils.ServiceConfig
}

func (s *ClientConfig) Validate() error {
	if s.ServerPort <= 0 || s.ServerPort > 65535 {
		return fmt.Errorf("error:Invalid Server port")
	}
	if s.ServerHost == "" {
		return fmt.Errorf("error:ServerHost cannot be empty")
	}
	if s.ServiceConfig == nil {
		return fmt.Errorf("error:ServiceConfig cannot be nil")
	}
	return nil
}

func (s *ClientConfig) String() string {
	return fmt.Sprintf("ClientConfig{ServerHost: %s, ServerPort: %d, ServiceConfig: %v}", s.ServerHost, s.ServerPort, s.ServiceConfig)
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
