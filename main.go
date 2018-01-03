package main

import (
	"github.com/painterdrown/virtual-routing/router"
	"github.com/painterdrown/virtual-routing/util"
)

func main() {
	defer util.HandleErr()
	router.Init()
	go router.Listen()
	go router.BroadcastPeriodically()
	go router.UpdateRoutingTablePeriodically()
	router.RunCmd()

	// router.Test()
}
