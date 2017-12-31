package router

import "fmt"

const bigenough = 9999

var name string     // name 是该主机的名字
var port int        // port 用来标识不同的主机
var ready = false   // ready 为 true 表示该主机已经配置完毕
var updated = false // updated 表示路由表是否需要更新

var lock1 = make(chan bool, 1)
var lock2 = make(chan bool, 1)

var all = make(map[int]bool)
var near = make(map[int]bool)
var dist = make(map[int]int)
var prev = make(map[int]int)
var cost = make(map[int]map[int]int)   // cost 相当于二维数组，用来记录每条链路上面的花费
var broadcasted = make(map[int64]bool) // Broadcasted 储存已经转发的广播信息 ID

func init() {
	all[port] = true
}

// ShowCost .
func ShowCost() {
	for k1, v1 := range cost {
		fmt.Printf("[cost][%d]: ", k1)
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)", k2, v2)
		}
		fmt.Printf("\n")
	}
}

// ShowDist .
func ShowDist() {
	fmt.Printf("[dist][%d]: ", port)
	for k, v := range dist {
		fmt.Printf("(%d,%d)", k, v)
	}
	fmt.Printf("\n")
}

// ShowPrev .
func ShowPrev() {
	fmt.Printf("[prev][%d]: ", port)
	for k, v := range prev {
		fmt.Printf("(%d,%d)", k, v)
	}
	fmt.Printf("\n")
}

// ShowNear .
func ShowNear() {
	fmt.Printf("[near][%d]: ", port)
	for k := range near {
		fmt.Printf("%d ", k)
	}
	fmt.Printf("\n")
}
