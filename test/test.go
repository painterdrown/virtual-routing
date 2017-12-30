package test

import (
	"github.com/painterdrown/virtual-routing/global"
)

// SetTestData .
func SetTestData() {
	global.Port = 1

	global.Cost[1] = make(map[int]int)
	global.Cost[2] = make(map[int]int)
	global.Cost[3] = make(map[int]int)
	global.Cost[4] = make(map[int]int)
	global.Cost[5] = make(map[int]int)
	global.Cost[1][1] = 0
	global.Cost[1][2] = global.INFINITE
	global.Cost[1][3] = 1
	global.Cost[1][4] = 4
	global.Cost[1][5] = 6
	global.Cost[2][1] = global.INFINITE
	global.Cost[2][2] = 0
	global.Cost[2][3] = global.INFINITE
	global.Cost[2][4] = 3
	global.Cost[2][5] = 2
	global.Cost[3][1] = 1
	global.Cost[3][2] = global.INFINITE
	global.Cost[3][3] = 0
	global.Cost[3][4] = 2
	global.Cost[3][5] = global.INFINITE
	global.Cost[4][1] = 4
	global.Cost[4][2] = 3
	global.Cost[4][3] = 2
	global.Cost[4][4] = 0
	global.Cost[4][5] = global.INFINITE
	global.Cost[5][1] = 6
	global.Cost[5][2] = 2
	global.Cost[5][3] = global.INFINITE
	global.Cost[5][4] = global.INFINITE
	global.Cost[5][5] = 0

	global.All[1] = true
	global.All[2] = true
	global.All[3] = true
	global.All[4] = true
	global.All[5] = true

	global.Near[3] = true
	global.Near[4] = true
	global.Near[5] = true

	global.Ready = true
}
