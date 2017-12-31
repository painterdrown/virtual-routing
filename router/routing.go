package router

import (
	"strconv"
	"strings"
	"time"
)

// UpdateRoutingTablePeriodically 开始周期性地更新路由表。
func UpdateRoutingTablePeriodically() {
	const interval = 60 * time.Second
	ticker := time.NewTicker(interval)
	for _ = range ticker.C {
		if ready && updated {
			updateRoutingTable()
			updated = false
			ShowNear() // DEBUG
			ShowDist() // DEBUG
			ShowPrev() // DEBUG
			ShowCost() // DEBUG
		}
	}
}

func updateCost(source int, costs []string) {
	all[source] = true
	if cost[source] == nil {
		cost[source] = make(map[int]int)
	}
	for _, v := range costs {
		parts := strings.Split(v, " ")
		dest, _ := strconv.Atoi(parts[0])
		c, _ := strconv.Atoi(parts[1])
		all[dest] = true
		if dest == port {
			near[source] = true
		}
		if cost[dest] == nil {
			cost[dest] = make(map[int]int)
		}
		if cost[source][dest] != c {
			updated = true
			cost[source][dest] = c
			cost[dest][source] = c
		}
	}
}

func updateRoutingTable() {
	// 初始化
	var reached = make(map[int]bool)
	if dist == nil {
		dist = make(map[int]int)
	}
	for u := range all {
		if near[u] {
			dist[u] = cost[port][u]
		} else {
			dist[u] = bigenough
		}
	}
	dist[port] = 0
	prev[port] = port
	reached[port] = true

	numOfAll := len(all)
	for {
		min := bigenough
		var m, n int // n 是最近的主机, m 是到达 n 的上一台主机
		for u := range reached {
			var base = dist[u]
			for v := range all {
				if !reached[v] && base+cost[u][v] < min {
					min = base + cost[u][v]
					n = v
					m = u
				}
			}
		}
		// 找到 m, n
		reached[n] = true
		dist[n] = min
		prev[n] = m

		// 判断是否到达所有主机
		if len(reached) == numOfAll {
			break
		}
	}
}
