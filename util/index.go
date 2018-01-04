package util

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/painterdrown/virtual-routing/global"
)

// CheckErr .
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// WaitShortly .
func WaitShortly() {
	n := rand.Int63n(1000)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// Prompt .
func Prompt(format string, a ...interface{}) {
	if global.WatingForCmd {
		fmt.Printf("\n")
	}
	fmt.Printf(format+"\n> ", a...)
}

// HandleErr .
func HandleErr() {
	if e := recover(); e != nil {
		if msg, ok := e.(string); ok {
			Prompt(msg)
			Log(msg)
		} else if err, ok := e.(error); ok {
			msg = err.Error()
			Prompt(msg)
			Log(msg)
		} else {
			msg = "不知道为什么就崩溃了..."
			Prompt(msg)
			Log(msg)
		}
		os.Exit(1)
	}
}
