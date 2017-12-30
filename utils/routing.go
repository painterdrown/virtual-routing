package utils

import (
	"strconv"
	"strings"
	"time"

	"github.com/painterdrown/virtual-routing/global"
)

// UpdateCost .
func UpdateCost(source int, costs []string) {
	if global.Cost[source] == nil {
		global.Cost[source] = make(map[int]int)
	}
	for _, v := range costs {
		parts := strings.Split(v, " ")
		dest, _ := strconv.Atoi(parts[0])
		cost, _ := strconv.Atoi(parts[1])
		if dest == global.Port {
			global.Near[source] = true
		}
		if global.Cost[dest] == nil {
			global.Cost[dest] = make(map[int]int)
		}
		if global.Cost[source][dest] != cost {
			global.Updated = true
			global.Cost[source][dest] = cost
			global.Cost[dest][source] = cost
		}
	}
	global.Updated = false
}

// UpdateRoutingTable .
func UpdateRoutingTable() {
	// 初始化
	var reached = make(map[int]bool)
	if global.Dist == nil {
		global.Dist = make(map[int]int)
	}
	for u := range global.All {
		if global.Near[u] {
			global.Dist[u] = global.Cost[global.Port][u]
		} else {
			global.Dist[u] = global.INFINITE
		}
	}
	global.Dist[global.Port] = 0
	global.Prev[global.Port] = global.Port
	reached[global.Port] = true

	numOfAll := len(global.All)
	for {
		min := global.INFINITE
		var m, n int // n 是最近的主机, m 是到达 n 的上一台主机
		for u := range reached {
			var base = global.Dist[u]
			for v := range global.All {
				if !reached[v] && base+global.Cost[u][v] < min {
					min = base + global.Cost[u][v]
					n = v
					m = u
				}
			}
		}
		// 找到 m, n
		reached[n] = true
		global.Dist[n] = min
		global.Prev[n] = m

		// 判断是否到达所有主机
		if len(reached) == numOfAll {
			break
		}
	}
}

// UpdateRoutingTablePeriodically .
func UpdateRoutingTablePeriodically() {
	const interval = 60 * time.Second
	ticker := time.NewTicker(interval)
	for _ = range ticker.C {
		if global.Updated {
			UpdateRoutingTable()
			global.Updated = false
			global.ShowDist() // DEBUG
			global.ShowPrev() // DEBUG
		}
	}
}
