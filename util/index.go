package util

import (
	"fmt"
	"math/rand"
	"time"
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
}
