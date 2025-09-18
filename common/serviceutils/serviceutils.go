package serviceutils

import (
	logger "github.com/eyanshu1997/tunnlrx/common/log"
)

// This file contains utility functions for services like reading config files
// initializing common services such as logger

var (
	Log *logger.CustomLogger
)

func InitServiceUtils(serviceConfig *ServiceConfig) {
	var err error
	Log, err = logger.InitLogger(serviceConfig.Log.Level, serviceConfig.Log.FilePath, serviceConfig.Log.IncludeStdio)
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

}
