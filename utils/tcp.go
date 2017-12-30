package utils

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/painterdrown/virtual-routing/global"
)

// Listen .
func Listen() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(global.Port))
	CheckErr(err)
	for {
		conn, err := ln.Accept()
		CheckErr(err)
		go HandleConn(conn)
	}
}

// HandleConn .
func HandleConn(conn net.Conn) {
	defer conn.Close()
	var buffer [512]byte
	for {
		n, err := conn.Read(buffer[0:])
		if err != nil {
			if err.Error() == "EOF" {
				return
			}
			panic(err)
		}
		msg := string(buffer[0:n])
		Prompt("Recieving: " + msg)
		HandleMsg(msg)
	}
}

// HandleMsg .
func HandleMsg(msg string) {
	if len(msg) == 0 {
		return
	}
	parts := strings.Split(msg, "|")

	// 如果是广播信息
	if parts[0] == "B" {
		bid, _ := strconv.ParseInt(parts[1], 10, 64)
		source, _ := strconv.Atoi(parts[2])

		// 判断该广播信息是否已经被该主机广播过
		if global.Broadcasted[bid] {
			return
		}
		global.Broadcasted[bid] = true

		// 更新 Cost
		if UpdateCost(source, parts[3:]) && global.Ready {
			// UpdateRoutingTable()
			global.ShowCost() // DEBUG
		}

		// 向其他路由器继续转发
		Broadcast(msg)
	}

	// 如果是路由信息
	if parts[0] == "R" {

	}
}

// Communicate .
func Communicate(port int, msg string) {
	conn, err := net.Dial("tcp", "0.0.0.0:"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, msg)
	conn.Close()
}

// Broadcast .
func Broadcast(msg string) {
	for port := range global.Near {
		if port != global.Port {
			Communicate(port, msg)
		}
	}
	Prompt("Broadcasting: " + msg)
}

// BroadcastPeriodically .
func BroadcastPeriodically() {
	const interval = 8 * time.Second
	ticker := time.NewTicker(interval)
	for _ = range ticker.C {
		msg := GenerateBroadcastMsg()
		Broadcast(msg)
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
