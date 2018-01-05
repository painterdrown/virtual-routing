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
	case "port":
		p, _ := strconv.Atoi(args[1])
		if port != -1 {
			util.Prompt("错误: 不能重复配置端口")
		}
		if testListen(p) {
			port = p
			util.InitLogger(port, 2)
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
		ready = true
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
			util.Prompt("错误: 发送目标不能是自己")
			break
		}
		qry := "Q|" + strconv.Itoa(port) + "|" + strconv.Itoa(p)
		send(controller, qry)
		msg := "R|" + strconv.Itoa(port) + "|" + args[1] + "|" + args[2]
		n := <-next // 等待请求结果
		send(n, msg)
		break
	case "exit":
		msg := "D|" + strconv.Itoa(port)
		if port == controller {
			for u := range all {
				send(u, msg)
			}
		} else {
			send(controller, msg)
		}
		os.Exit(0)
	default:
		util.Prompt("错误: 无效的命令")
		break
	}
}
