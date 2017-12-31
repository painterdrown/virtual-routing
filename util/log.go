package util

import (
	"log"
	"os"
)

var logger *log.Logger

// Log .
func Log(format string, a ...interface{}) {
	if logger == nil {
		logFile, err := os.Create("info.log")
		CheckErr(err)
		defer logFile.Close()
		logger = log.New(logFile, "[test]", log.LstdFlags)
	}
	logger.Printf(format+"\n", a...)
}
