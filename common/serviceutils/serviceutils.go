package serviceutils

import logger "github.com/eyanshu1997/tunnlrx/common/log"

// This file contains utility functions for services like reading config files
// initializing common services such as logger

var (
	Log *logger.CustomLogger
)

func InitServiceUtils() {
	var err error
	Log, err = logger.InitLogger(logger.INFO, "service.log", true)
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
}
