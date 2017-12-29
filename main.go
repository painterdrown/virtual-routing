package main

import (
	"os"

	"github.com/painterdrown/virtual-routing/global"
	"github.com/painterdrown/virtual-routing/utils"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: $GOBIN/virtual-routing <port>")
		os.Exit(1)
	}
	global.Port = os.Args[1]
	go utils.Listen()
	utils.Config()
	go utils.BroadcastPeriodically()
}
