package router1

import (
	"sync"
)

const bigenough = 9999

var port = -1       // port 用来标识不同的主机
var ready = false   // ready 为 true 表示该主机已经配置完毕
var updated = false // updated 表示路由表是否需要更新

var lock1 = new(sync.Mutex)
var lock2 = new(sync.Mutex)
var lock3 = new(sync.Mutex)

var all = make(map[int]bool)
var near = make(map[int]bool)
var dist = make(map[int]int)
var prev = make(map[int]int)
var cost = make(map[int]map[int]int)   // cost 相当于二维数组，用来记录每条链路上面的花费
var broadcasted = make(map[int64]bool) // broadcasted 储存已经转发的广播信息 ID

func init() {}
