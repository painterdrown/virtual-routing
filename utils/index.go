package utils

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

// Prompt .
func Prompt(msg string) {
	fmt.Printf("%s\n", msg)
}

// WaitShortly .
func WaitShortly() {
	n := rand.Int63n(1000)
	time.Sleep(time.Duration(n) * time.Millisecond)
}
