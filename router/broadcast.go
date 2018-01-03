package router

import (
	"strconv"
	"time"

	"github.com/painterdrown/virtual-routing/util"
)

// BroadcastPeriodically 周期性地广播自己的路由信息。
func BroadcastPeriodically() {
	const interval = 5 * time.Second
	ticker := time.NewTicker(interval)
	for _ = range ticker.C {
		if ready {
			msg := generateBroadcastMsg()
			broadcast(msg)
		}
	}
}

func broadcast(msg string) {
	for p := range near {
		if p != port {
			send(p, msg)
		}
	}
	util.Log("广播: " + msg)
}

func getTimestamp() int64 {
	return time.Now().UnixNano()
}

func generateBroadcastMsg() string {
	bid := getTimestamp()
	broadcasted[bid] = true
	msg := "B|" + strconv.FormatInt(bid, 10) + "|" + strconv.Itoa(port)
	for p := range near {
		msg += "|" + strconv.Itoa(p) + " " + strconv.Itoa(cost[port][p])
	}
	return msg
}
