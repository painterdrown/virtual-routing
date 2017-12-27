package utils

import (
	"fmt"

	"github.com/painterdrown/virtual-routing/global"
)

// RecieveCmd .
func RecieveCmd() {
	for {
		var cmd string
		fmt.Printf("[%s] > ", global.Host)
		fmt.Scanf("%s", &cmd)
	}
}

func handleCmd(cmd string) {
	println(cmd)
}
