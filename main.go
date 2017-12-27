package main

import (
	"os"

	"github.com/painterdrown/virtual-routing/global"
	"github.com/painterdrown/virtual-routing/utils"
)

func main() {
	if len(os.Args) < 3 {
		println("Usage: virtual-routing <ip> <port>")
		return
	}
	global.Host = os.Args[1]
	global.Port = os.Args[2]
	go utils.Listen(global.Port)
	utils.RecieveCmd()
}
