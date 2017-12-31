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
	fmt.Printf(format+"\n", a...)
	if global.WatingForCmd {
		fmt.Printf("> ")
	}
}

// HandleErr .
func HandleErr() {
	if e := recover(); e != nil {
		if msg, ok := e.(string); ok {
			Log(msg)
		} else if err, ok := e.(error); ok {
			Log(err.Error())
		} else {
			Log("不知道为什么我就崩溃了...")
		}
		os.Exit(1)
	}
}
