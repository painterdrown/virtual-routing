package router2

import (
	"os"
	"strconv"
	"sync"

	"github.com/painterdrown/virtual-routing/util"
)

const bigenough = 9999

var mode = 1 // 1 是 LS 自治, 2 是 LS 中控, 3 是 DV
var controller = -1
var port = -1       // port 用来标识不同的主机
var name string     // name 是该主机的名字
var ready = false   // ready 为 true 表示该主机已经配置完毕
var updated = false // updated 表示路由表是否需要更新

var lock1 = new(sync.Mutex)
var lock2 = new(sync.Mutex)

var all = make(map[int]bool)
var near = make(map[int]bool)
var dist = make(map[int]int)
var prev = make(map[int]int)
var cost = make(map[int]map[int]int)   // cost 相当于二维数组，用来记录每条链路上面的花费
var broadcasted = make(map[int64]bool) // Broadcasted 储存已经转发的广播信息 ID

func init() {}

// Init 初始化路由的基本信息。
func Init() {
	if len(os.Args) < 3 {
		panic("缺乏参数 port 或 name")
	}
	p, err := strconv.Atoi(os.Args[1])
	util.CheckErr(err)
	if testPort(p) {
		port = p
	} else {
		panic("监听端口 " + strconv.Itoa(p) + " 出错，或者该端口已被占用。请选择其他端口！")
	}
	name = os.Args[2]
	util.InitLogger(name, port)
}
