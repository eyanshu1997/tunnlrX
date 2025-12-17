package log

//This file replaces the original logger for go lang with a custom logger
//that prints logs in a structured format with timestamps and log levels, it also prints the calling filename and line number and (function).
//It also supports different log levels like DEBUG, INFO, ERROR, PANIC.

//Typical log would look like this:
//2023-10-10 10:10:10 [INFO] main.go:10 (main.main) - Server started on port 8080

// typical logging priority would be
// Panic
// Error
// Info
// Debug

// the functions would be panic, error, info, debug for similicity but all support formatted strings like fmt.Printf

// ideally logger should be initialized once in main.go and then used throughout the application and passed around as global in app
// but we can assume there is always one logger instance for simplicity
// it should be responsibility of the init service for the function to have that logger instance initialization

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type logLevel int32

const (
	PANIC logLevel = iota
	ERROR
	INFO
	DEBUG
)

var logLevelNames = map[logLevel]string{
	PANIC: "PANIC",
	ERROR: "ERROR",
	INFO:  "INFO",
	DEBUG: "DEBUG",
}
var logLevelValues = map[string]logLevel{
	"panic": PANIC,
	"error": ERROR,
	"info":  INFO,
	"debug": DEBUG,
}

type CustomLogger struct {
	level logLevel
}

var logger *CustomLogger

func InitLogger(logLevelStr string) error {
	var logLevel logLevel
	if lvl, ok := logLevelValues[logLevelStr]; ok {
		logLevel = lvl
	} else {
		return fmt.Errorf("invalid log level: %s", logLevelStr)
	}

	logger = &CustomLogger{
		level: logLevel,
	}
	return nil
}

func (l *CustomLogger) logf(level logLevel, format string, v ...interface{}) {
	if level > logger.level {
		// Below the current log level, do not log
		// for eg if current level is INFO, do not log DEBUG
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	// Get caller info
	_, file, line, ok := runtime.Caller(2)
	callerInfo := ""
	if ok {
		callerInfo = fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}
	prefix := fmt.Sprintf("%s [%s] %s - ", now, logLevelNames[level], callerInfo)
	log.Printf(prefix+format, v...)
}

func Panic(format string, v ...interface{}) {
	logger.logf(PANIC, format, v...)
	panic(fmt.Sprintf(format, v...))
}

func Error(format string, v ...interface{}) {
	logger.logf(ERROR, format, v...)
}

func Info(format string, v ...interface{}) {
	logger.logf(INFO, format, v...)
}

func Fatalf(format string, v ...interface{}) {
	logger.logf(ERROR, format, v...)
	os.Exit(1)
}

func Debug(format string, v ...interface{}) {
	logger.logf(DEBUG, format, v...)
}
