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
		Prompt(msg)
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
		source, _ := strconv.Atoi(parts[1])

		// 更新 Cost
		if UpdateCost(source, parts[2:]) && global.Ready {
			UpdateRoutingTable()
		}

		// 向其他路由器继续转发
		for port := range global.DC {
			// 防止形成循环
			if port == source || !global.All[port] {
				continue
			}
			Communicate(port, msg)
		}
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
func Broadcast() {
	msg := "B|" + strconv.Itoa(global.Port)
	for port := range global.DC {
		conn, err := net.Dial("tcp", "0.0.0.0:"+strconv.Itoa(port))
		if err != nil {
			global.All[port] = false
			global.Cost[global.Port][port] = global.INFINITE
		} else {
			msg += "|" + strconv.Itoa(port) + " " + strconv.Itoa(global.Cost[global.Port][port])
		}
		conn.Close()
	}
	for port := range global.DC {
		if global.All[port] {
			Communicate(port, msg)
		}
	}
}

// BroadcastPeriodically .
func BroadcastPeriodically() {
	const interval = 30 * time.Second
	ticker := time.NewTicker(interval)
	for _ = range ticker.C {
		Broadcast()
	}
}
