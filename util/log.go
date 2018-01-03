package util

import (
	"log"
	"os"
	"strconv"
)

var logger *log.Logger

// InitLogger .
func InitLogger(name string, port int) {
	logFile, err := os.Create("logs/" + strconv.Itoa(port) + ".log")
	CheckErr(err)
	logger = log.New(logFile, "["+name+"]", log.LstdFlags)
}

// Log .
func Log(format string, a ...interface{}) {
	if logger == nil {
		logFile, err := os.Create("logs/error.log")
		CheckErr(err)
		logger = log.New(logFile, "[ERROR]", log.LstdFlags)
	}
	logger.Printf(format+"\n", a...)
}
