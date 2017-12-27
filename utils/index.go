package utils

import (
	"fmt"

	"github.com/painterdrown/virtual-routing/global"
)

// CheckErr .
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Prompt .
func Prompt(msg string) {
	fmt.Printf("\n%s\n[%s] > ", msg, global.Host)
}
