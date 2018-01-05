package router2

import (
	"strconv"
	"strings"
	"time"
)

func answer(u, v int) int {
	var res int
	for {
		if prev[u][v] == u {
			res = v
			break
		} else {
			v = prev[u][v]
		}
	}
	return res
}

func updateRoutingTablePeriodically() {
	const interval = 22 * time.Second
	ticker := time.NewTicker(interval)
	for _ = range ticker.C {
		if ready && updated {
			updateRoutingTable()
			updated = false
		}
	}
}

func updateRoutingTable() {
	for p := range all {
		if dist[p] == nil {
			dist[p] = make(map[int]int)
		}
		if prev[p] == nil {
			prev[p] = make(map[int]int)
		}
		for u := range all {
			if near[p][u] {
				dist[p][u] = cost[p][u]
				prev[p][u] = p
			} else {
				dist[p][u] = bigenough
			}
		}
		// 初始化
		dist[p][p] = 0
		prev[p][p] = p
		var reached = make(map[int]bool)
		reached[p] = true
		// 开始计算最短路径
		numOfAll := len(all)
		for {
			min := bigenough
			var w int // n 是最近的主机, m 是到达 n 的上一台主机
			for u := range dist[p] {
				if !reached[u] && dist[p][u] < min {
					min = dist[p][u]
					w = u
				}
			}
			// w 就是找到的最小开销的路由器
			reached[w] = true
			for v := range cost[w] {
				if !reached[v] && dist[p][w]+cost[w][v] < dist[p][v] {
					dist[p][v] = dist[p][w] + cost[w][v]
					prev[p][v] = w
				}
			}
			// 判断是否到达所有主机
			if len(reached) == numOfAll {
				break
			}
		}
	}
}

func updateCost(source int, costs []string) {
	all[source] = true
	down[source] = false
	if near[source] == nil {
		near[source] = make(map[int]bool)
	}
	if cost[source] == nil {
		cost[source] = make(map[int]int)
	}
	for _, v := range costs {
		parts := strings.Split(v, " ")
		dest, _ := strconv.Atoi(parts[0])
		c, _ := strconv.Atoi(parts[1])
		if _, ok := down[dest]; ok {
			if down[dest] {
				continue
			}
		} else {
			down[dest] = false
		}
		if near[dest] == nil {
			near[dest] = make(map[int]bool)
		}
		if cost[dest] == nil {
			cost[dest] = make(map[int]int)
		}
		all[dest] = true
		near[source][dest] = true
		near[dest][source] = true
		if cost[source][dest] != c {
			cost[source][dest] = c
			cost[dest][source] = c
			updated = true
		}
	}
}
