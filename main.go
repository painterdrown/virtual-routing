package main

import (
	"os"

	"github.com/painterdrown/virtual-routing/global"
	"github.com/painterdrown/virtual-routing/router"
	"github.com/painterdrown/virtual-routing/util"
)

func main() {
	defer util.HandleErr()
	router.Init()
	go router.Listen()
	router.Config()
	go router.BroadcastPeriodically()
	go router.UpdateRoutingTablePeriodically()
	os.Exit(<-global.Exit)
}
