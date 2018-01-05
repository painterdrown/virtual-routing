package router2

import (
	"sync"
)

const bigenough = 9999

var controller = -1
var port = -1 // port 用来标识不同的主机
var next = make(chan int, 1)
var ready = false   // ready 为 true 表示该主机已经配置完毕
var updated = false // updated 表示路由表是否需要更新

var lock = new(sync.Mutex)

// controller
var all = make(map[int]bool)
var down = make(map[int]bool)
var near = make(map[int]map[int]bool)
var dist = make(map[int]map[int]int)
var prev = make(map[int]map[int]int)
var cost = make(map[int]map[int]int) // cost 相当于二维数组，用来记录每条链路上面的花费

// slave
var snear = make(map[int]bool)
var scost = make(map[int]int)

func init() {}
