package utils

import (
	"fmt"
	"strconv"

	"github.com/painterdrown/virtual-routing/global"
)

// Config .
func Config() {
	args := make([]string, 3)
	for {
		fmt.Printf("[%d] > ", global.Port)
		fmt.Scanf("%s %s %s", &args[0], &args[1], &args[2])
		if !HandleCmd(args) {
			break
		}
		args[0] = ""
		args[1] = ""
		args[2] = ""
	}
}

// HandleCmd .
func HandleCmd(args []string) bool {
	if args[0] == "" {
		return true
	}

	// 配置主机名称
	if args[0] == "name" {
		global.Name = args[1]
	}

	// 配置与主机相连的拓扑以及花费
	if args[0] == "connect" {
		port, _ := strconv.Atoi(args[1])
		cost, _ := strconv.Atoi(args[2])
		if global.Cost[global.Port] == nil {
			global.Cost[global.Port] = make(map[int]int)
		}
		if global.Cost[port] == nil {
			global.Cost[port] = make(map[int]int)
		}
		global.All[port] = true
		global.Near[port] = true
		global.Cost[global.Port][port] = cost
		global.Cost[port][global.Port] = cost
		return true
	}

	// 完成配置
	if args[0] == "ok" {
		global.Ready = true
		return false
	}

	return true
}
