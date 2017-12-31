package router

import (
	"strconv"
	"time"

	"github.com/painterdrown/virtual-routing/util"
)

// BroadcastPeriodically .
func BroadcastPeriodically() {
	const interval = 30 * time.Second
	ticker := time.NewTicker(interval)
	for _ = range ticker.C {
		if ready {
			msg := generateBroadcastMsg()
			broadcast(msg, -1)
		}
	}
}

func broadcast(msg string, except int) {
	for p := range near {
		if p != port && p != except {
			send(p, msg)
		}
	}
	util.Log("广播: " + msg)
}

func generateBroadcastMsg() string {
	bid := time.Now().UnixNano()
	broadcasted[bid] = true
	msg := "B|" + strconv.FormatInt(bid, 10) + "|" + strconv.Itoa(port)
	for p := range near {
		msg += "|" + strconv.Itoa(p) + " " + strconv.Itoa(cost[port][p])
	}
	return msg
}
