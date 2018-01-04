package router2

import (
	"strconv"
	"time"
)

func reportNeighborsPeriodically() {
	const interval = 11 * time.Second
	ticker := time.NewTicker(interval)
	for _ = range ticker.C {
		if ready {
			reportNeighbors()
		}
	}
}

func reportNeighbors() {
	if len(snear) == 0 {
		return
	}
	msg := "N|" + strconv.Itoa(port)
	for n := range snear {
		msg += "|" + strconv.Itoa(n) + " " + strconv.Itoa(scost[n])
	}
	send(controller, msg)
}

func connect(p, c int) {
	if cost[port] == nil {
		cost[port] = make(map[int]int)
	}
	snear[p] = true
	scost[p] = c
}

func query(u, v int) {
	msg := "Q|" + strconv.Itoa(u) + "|" + strconv.Itoa(v)
	send(controller, msg)
}
