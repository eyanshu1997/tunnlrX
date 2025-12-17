package serviceutils

import (
	"context"

	logger "github.com/eyanshu1997/tunnlrx/common/log"
)

// This file contains utility functions for services like reading config files
// initializing common services such as logger

var (
	Log *logger.CustomLogger
)

func InitServiceUtils(serviceConfig *ServiceConfig) {
	var err error
	Log, err = logger.InitLogger(serviceConfig.Log.Level)
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	Log.Debug("Config %+v", serviceConfig)

}

func GetContextWithMetadata() context.Context {
	// Add any common metadata to context if needed
	return context.Background()
}
