package global

// INFINITE .
const INFINITE = int(^uint(0) >> 1)

// Name 是该主机的名字
var Name string

// Port 用来标识不同的主机
var Port string

// Cost 相当于二维数组，用来记录每条链路上面的花费
var Cost = make(map[string]map[string]int)

// Ready 为 true 表示该主机已经配置完毕
var Ready = false

// DC 是直接与该主机相连的其他主机列表
var DC = make([]string, 0)

// DCUp 表示 DC 中的主机是否在线上
var DCUp = make([]bool, 0)
