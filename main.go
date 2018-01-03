package main

import (
	"os"
	"strconv"

	"github.com/painterdrown/virtual-routing/router1"
	"github.com/painterdrown/virtual-routing/util"
)

func main() {
	defer util.HandleErr()
	if len(os.Args) < 2 {
		panic("错误: 缺乏参数 mode")
	}
	mode, _ := strconv.Atoi(os.Args[1])
	switch mode {
	case 1:
		router1.RunCmd()
		break
	case 2:
		break
	case 3:
		break
	default:
		panic("错误: 参数 mode 的值只能是 1/2/3")
	}
}
