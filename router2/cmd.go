package router2

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
	case "controller":
		p, _ := strconv.Atoi(args[1])
		controller = p
		break
	case "connect":
		p, _ := strconv.Atoi(args[1])
		c, _ := strconv.Atoi(args[2])
		connect(p, c)
		break
	case "ok":
		// 已经 ok 过一次
		if ready {
			break
		}
		ready = true
		go listen()
		if controller == port {
			go updateRoutingTablePeriodically()
		} else {
			go reportNeighborsPeriodically()
		}
		util.Prompt("配置完成，正在监听 %d 端口...", port)
		break
	case "info":
		if controller != port {
			util.Prompt("只有 controller 能查看信息")
			break
		}
		showInfo()
		break
	case "send":
		p, _ := strconv.Atoi(args[1])
		if port == controller {
			send(p, args[2])
			break
		}
		if p == port {
			util.Prompt("错误: 发送的目标不能是自己")
			break
		}
		msg := "R|" + strconv.Itoa(port) + "|" + args[1] + "|" + args[2]
		send(<-next, msg)
		break
	case "exit":
		msg := "D|" + strconv.Itoa(port)
		send(controller, msg)
		os.Exit(0)
	default:
	}
}
