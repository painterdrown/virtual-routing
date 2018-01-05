package router3

import (
	"strconv"
	"strings"
)

func updateRoutingTable() {
}

func updateCost(source int, costs []string) {
	if !near[source] {
		all[source] = true
		near[source] = true
		next[source] = source
		for _, v := range costs {
			parts := strings.Split(v, " ")
			dest, _ := strconv.Atoi(parts[0])
			c, _ := strconv.Atoi(parts[1])
			if dest == port {
				cost[source] = c
				dist[source] = c
				updated = true
				break
			}
		}
	}
	for _, v := range costs {
		parts := strings.Split(v, " ")
		dest, _ := strconv.Atoi(parts[0])
		c, _ := strconv.Atoi(parts[1])
		all[dest] = true
		if _, ok := dist[dest]; !ok || dist[source]+c < dist[dest] {
			dist[dest] = dist[source] + c
			next[dest] = source
			if next[source] != port {
				next[dest] = next[source]
			}
			updated = true
		}
	}
}

func connect(p, c int) {
	all[p] = true
	near[p] = true
	next[p] = p
	cost[p] = c
	dist[p] = cost[p]
}

func shareDist() {
	if len(near) == 0 {
		return
	}
	msg := "S|" + strconv.Itoa(port)
	for k, v := range dist {
		msg += "|" + strconv.Itoa(k) + " " + strconv.Itoa(v)
	}
	tellNeighbors(msg)
}

func tellNeighbors(msg string) {
	for nb := range near {
		send(nb, msg)
	}
}
