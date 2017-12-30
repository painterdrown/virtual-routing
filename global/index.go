package global

// INFINITE .
// const INFINITE = int(^uint(0) >> 1)
const INFINITE = 9999

// Name 是该主机的名字
var Name string

// Port 用来标识不同的主机
var Port int

// Ready 为 true 表示该主机已经配置完毕
var Ready = false

// Cost 相当于二维数组，用来记录每条链路上面的花费
var Cost = make(map[int]map[int]int)

// Dist .
var Dist = make(map[int]int)

// Prev .
var Prev = make(map[int]int)

// Near 是直接与该主机相连的其他主机列表
var Near = make(map[int]bool)

// All 表示 网络中所有主机的在线情况
var All = make(map[int]bool)

// Broadcasted 储存已经转发的广播信息 ID
var Broadcasted = make(map[int64]bool)

// Exit .
var Exit = make(chan bool)

func init() {
	All[Port] = true
}
