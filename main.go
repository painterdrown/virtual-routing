package main

import (
	"os"
	"strconv"

	"github.com/painterdrown/virtual-routing/global"
	"github.com/painterdrown/virtual-routing/router"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: $GOBIN/virtual-routing <port>")
		global.Exit <- 1
	}
	port, _ := strconv.Atoi(os.Args[1])
	go router.Listen()
	router.Config(port)
	go router.BroadcastPeriodically()
	go router.UpdateRoutingTablePeriodically()
	os.Exit(<-global.Exit)
}
