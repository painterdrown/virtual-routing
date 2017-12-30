package utils

import (
	"fmt"
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
