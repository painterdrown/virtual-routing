package router3

import (
	"sync"
)

const bigenough = 9999

var port = -1       // port 用来标识不同的主机
var ready = false   // ready 为 true 表示该主机已经配置完毕
var updated = false // updated 表示路由表是否需要更新

var lock1 = new(sync.Mutex)
var lock2 = new(sync.Mutex)

var all = make(map[int]bool)
var near = make(map[int]bool)
var cost = make(map[int]int)
var dist = make(map[int]int)
var next = make(map[int]int)

func init() {}
