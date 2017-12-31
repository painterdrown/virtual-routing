package router

import (
	"fmt"
	"strconv"

	"github.com/painterdrown/virtual-routing/util"
)

// RunCmd 接收用户的命令。
func RunCmd() {
	args := make([]string, 4)
	for {
		fmt.Printf("> ")
		waitingForCmd = true
		fmt.Scanf("%s %s %s", &args[0], &args[1], &args[2], &args[3])
		waitingForCmd = false
		if !handleCmd(args) {
			break
		}
		args[0] = ""
		args[1] = ""
		args[2] = ""
		args[3] = ""
	}
}

func handleCmd(args []string) bool {
	if len(args) == 0 {
		return true
	}

	// 完成配置
	if args[0] == "ok" {

	}

	switch args[0] {
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
	}

	return true
}
