package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/painterdrown/virtual-routing/global"
)

// Config .
func Config() {
	for {
		var cmd string
		fmt.Printf("[%s] > ", global.Port)
		fmt.Scanf("%s", &cmd)
		if !HandleCmd(cmd) {
			break
		}
	}
}

// HandleCmd .
func HandleCmd(cmd string) bool {
	if len(cmd) == 0 {
		return true
	}
	parts := strings.Split(cmd, " ")

	// 配置主机名称
	if parts[0] == "name" {
		global.Name = parts[1]
	}

	// 配置与主机相连的拓扑以及花费
	if parts[0] == "connect" {
		port := parts[1]
		cost, _ := strconv.Atoi(parts[2])
		global.Cost[global.Port][port] = cost
		return true
	}

	// 完成配置
	if parts[0] == "ok" {
		global.Ready = true
		UpdateRoutingTable()
		return false
	}

	return true
}
