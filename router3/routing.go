package router3

import (
	"strconv"
	"strings"
)

func updateRoutingTable() {
}

func updateCost(source int, costs []string) {
	all[source] = true
	near[source] = true
	for _, v := range costs {
		parts := strings.Split(v, " ")
		dest, _ := strconv.Atoi(parts[0])
		c, _ := strconv.Atoi(parts[1])
		if _, ok := dist[dest]; !ok || dist[source]+c < dist[dest] {
			dist[dest] = dist[source] + c
			next[dest] = source
			updated = true
		}
	}
}

func connect(p, c int) {
	all[p] = true
	near[p] = true
	cost[p] = c
	dist[p] = cost[p]
}

func shareDist() {
	if len(near) == 0 {
		return
	}
	msg := "S|" + strconv.Itoa(port)
	for k, v := range cost {
		msg += "|" + strconv.Itoa(k) + " " + strconv.Itoa(v)
	}
	tellNeighbors(msg)
}

func tellNeighbors(msg string) {
	for nb := range near {
		send(nb, near)
	}
}
