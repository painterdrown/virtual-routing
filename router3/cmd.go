package router3

import (
	"fmt"
	"os"
	"strconv"

	"github.com/painterdrown/virtual-routing/global"
	"github.com/painterdrown/virtual-routing/util"
)

// RunCmd 接收用户的命令。
func RunCmd() {
	args := make([]string, 3)
	for {
		fmt.Printf("> ")
		global.WatingForCmd = true
		fmt.Scanf("%s %s %s", &args[0], &args[1], &args[2])
		global.WatingForCmd = false
		handleCmd(args)
		args[0] = ""
		args[1] = ""
		args[2] = ""
	}
}

func handleCmd(args []string) {
	if len(args) == 0 {
		return
	}
	op := args[0]
	switch op {
	case "port":
		p, _ := strconv.Atoi(args[1])
		if port != -1 {
			util.Prompt("错误: 不能重复配置端口")
		}
		if testListen(p) {
			port = p
			all[port] = true
			cost[port] = 0
			util.InitLogger(port, 3)
			go listen()
		} else {
			util.Prompt("错误: 该端口已被占用, 请选择其他端口")
		}
		break
	case "connect":
		p, _ := strconv.Atoi(args[1])
		c, _ := strconv.Atoi(args[2])
		connect(p, c)
		break
	case "ok":
		if ready {
			break
		}
		allNbIsReady := true
		for n := range near {
			if !testConnection(n) {
				util.Prompt("错误: 邻居 %d 未配置端口")
				allNbIsReady = false
			}
		}
		if !allNbIsReady {
			break
		}
		ready = true
		shareDist()
		util.Prompt("配置完成，正在监听 %d 端口...", port)
		break
	case "info":
		showInfo()
		break
	case "send":
		p, _ := strconv.Atoi(args[1])
		if p == port {
			util.Prompt("错误: 发送目标不能是自己")
			break
		}
		if !all[p] {
			util.Prompt("错误: 发送目标不存在于当前网络中")
			break
		}
		msg := "R|" + strconv.Itoa(port) + "|" + strconv.Itoa(p) + "|" + args[2]
		n := next[p]
		send(n, msg)
		break
	case "exit":
		msg := "D|" + strconv.Itoa(port)
		tellNeighbors(msg)
		os.Exit(0)
	default:
		util.Prompt("错误: 无效的命令")
		break
	}
}
