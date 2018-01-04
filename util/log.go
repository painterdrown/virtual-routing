package util

import (
	"log"
	"os"
	"strconv"
)

var logger *log.Logger

// InitLogger .
func InitLogger(p, mode int) {
	filename := "logs/" + strconv.Itoa(p) + "-" + strconv.Itoa(mode) + ".log"
	logFile, err := os.Create(filename)
	CheckErr(err)
	logger = log.New(logFile, "["+strconv.Itoa(p)+"]", log.LstdFlags)
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
