package utils

import (
	"strconv"
	"time"

	"github.com/painterdrown/virtual-routing/global"
)

// Broadcast .
func Broadcast(msg string, except int) {
	for port := range global.Near {
		if port != global.Port && port != except {
			Communicate(port, msg)
			WaitShortly()
		}
	}
	Prompt("Broadcasting: " + msg)
}

// BroadcastPeriodically .
func BroadcastPeriodically() {
	const interval = 30 * time.Second
	ticker := time.NewTicker(interval)
	for _ = range ticker.C {
		msg := GenerateBroadcastMsg()
		Broadcast(msg, -1)
	}
}

// GenerateBroadcastMsg .
func GenerateBroadcastMsg() string {
	bid := time.Now().UnixNano()
	global.Broadcasted[bid] = true
	msg := "B|" + strconv.FormatInt(bid, 10) + "|" + strconv.Itoa(global.Port)
	for port := range global.Near {
		msg += "|" + strconv.Itoa(port) + " " + strconv.Itoa(global.Cost[global.Port][port])
	}
	return msg
}
