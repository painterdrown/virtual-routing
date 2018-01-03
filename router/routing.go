package router

import (
	"strconv"
	"strings"
	"time"
)

// UpdateRoutingTablePeriodically 开始周期性地更新路由表。
func UpdateRoutingTablePeriodically() {
	const interval = 10 * time.Second
	ticker := time.NewTicker(interval)
	for _ = range ticker.C {
		if ready && updated {
			updateRoutingTable()
			updated = false
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
	for u := range all {
		if near[u] {
			dist[u] = cost[port][u]
			prev[u] = port
		} else {
			dist[u] = bigenough
		}
	}
	dist[port] = 0
	prev[port] = port
	var reached = make(map[int]bool)
	reached[port] = true

	numOfAll := len(all)
	for {
		min := bigenough
		var w int // n 是最近的主机, m 是到达 n 的上一台主机
		for u := range dist {
			if !reached[u] && dist[u] < min {
				min = dist[u]
				w = u
			}
		}
		// w 就是找到的最小开销的路由器
		reached[w] = true
		for v := range cost[w] {
			if !reached[v] && dist[w]+cost[w][v] < dist[v] {
				dist[v] = dist[w] + cost[w][v]
				prev[v] = w
			}
		}
		// 判断是否到达所有主机
		if len(reached) == numOfAll {
			break
		}
	}
}
