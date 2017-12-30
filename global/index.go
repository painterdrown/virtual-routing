package global

// INFINITE .
// const INFINITE = int(^uint(0) >> 1)
const INFINITE = 9999

// Name 是该主机的名字
var Name string

// Port 用来标识不同的主机
var Port int

// Cost 相当于二维数组，用来记录每条链路上面的花费
var Cost = make(map[int]map[int]int)

// Prev .
var Dist = make(map[int]int)

// Prev .
var Prev = make(map[int]int)

// Ready 为 true 表示该主机已经配置完毕
var Ready = false

// DC 是直接与该主机相连的其他主机列表
var DC = make(map[int]bool)

// All 表示 网络中所有主机的在线情况
var All = make(map[int]bool)

// TestData .
func TestData() {
	Port = 1

	Cost[1] = make(map[int]int)
	Cost[2] = make(map[int]int)
	Cost[3] = make(map[int]int)
	Cost[4] = make(map[int]int)
	Cost[5] = make(map[int]int)
	Cost[1][1] = 0
	Cost[1][2] = INFINITE
	Cost[1][3] = 1
	Cost[1][4] = 4
	Cost[1][5] = 6
	Cost[2][1] = INFINITE
	Cost[2][2] = 0
	Cost[2][3] = INFINITE
	Cost[2][4] = 3
	Cost[2][5] = 2
	Cost[3][1] = 1
	Cost[3][2] = INFINITE
	Cost[3][3] = 0
	Cost[3][4] = 2
	Cost[3][5] = INFINITE
	Cost[4][1] = 4
	Cost[4][2] = 3
	Cost[4][3] = 2
	Cost[4][4] = 0
	Cost[4][5] = INFINITE
	Cost[5][1] = 6
	Cost[5][2] = 2
	Cost[5][3] = INFINITE
	Cost[5][4] = INFINITE
	Cost[5][5] = 0

	if Cost[6] == nil {
		println("haha")
	}

	All[1] = true
	All[2] = true
	All[3] = true
	All[4] = true
	All[5] = true

	DC[1] = true
	DC[3] = true
	DC[4] = true
	DC[5] = true

	Ready = true
}
