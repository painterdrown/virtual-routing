package router

import (
	"fmt"
	"os"
	"strconv"

	"github.com/painterdrown/virtual-routing/global"
	"github.com/painterdrown/virtual-routing/util"
)

// RunCmd 接收用户的命令。
func RunCmd() {
	args := make([]string, 4)
	for {
		fmt.Printf("> ")
		global.WatingForCmd = true
		fmt.Scanf("%s %s %s", &args[0], &args[1], &args[2], &args[3])
		global.WatingForCmd = false
		handleCmd(args)
		args[0] = ""
		args[1] = ""
		args[2] = ""
		args[3] = ""
	}
}

func handleCmd(args []string) {
	if len(args) == 0 {
		return
	}
	op := args[0]
	switch op {
	case "name":
		name = args[1]
		break
	case "connect":
		p, _ := strconv.Atoi(args[1])
		c, _ := strconv.Atoi(args[2])
		if cost[port] == nil {
			cost[port] = make(map[int]int)
		}
		if cost[p] == nil {
			cost[p] = make(map[int]int)
		}
		all[p] = true
		near[p] = true
		cost[port][p] = c
		cost[p][port] = c
		break
	case "ok":
		ready = true
		util.Prompt("配置完成，正在监听 %d 端口...", port)
		break
	case "shutdown":
		did := getTimestamp()
		msg := "D|" + strconv.FormatInt(did, 10) + "|" + strconv.Itoa(port)
		broadcast(msg)
		os.Exit(0)
	case "info":
		ShowInfo()
		break
	case "test":
		updateRoutingTable()
		break
	default:
	}
}
