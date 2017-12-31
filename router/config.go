package router

import (
	"fmt"
	"strconv"

	"github.com/painterdrown/virtual-routing/util"
)

// Config 进行虚拟路由的配置。
func Config() {
	args := make([]string, 3)
	for {
		fmt.Printf("[%d] > ", port)
		fmt.Scanf("%s %s %s", &args[0], &args[1], &args[2])
		if !handleCmd(args) {
			break
		}
		args[0] = ""
		args[1] = ""
		args[2] = ""
	}
}

func handleCmd(args []string) bool {
	if len(args) == 0 {
		return true
	}

	// 配置端口
	if args[0] == "port" {

	}

	// 配置名称
	if args[0] == "name" {
		name = args[1]
	}

	// 配置拓扑以及花费
	if args[0] == "connect" {
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
	}

	// 完成配置
	if args[0] == "ok" {
		if port == -1 {
			util.Prompt("必须配置端口：port <port>")
			return true
		}
		ready = true
		util.Prompt("配置完成，正在监听 %d 端口...", port)
		return false
	}

	return true
}
