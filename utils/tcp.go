package utils

import (
	"fmt"
	"net"
	"strconv"
	"strings"

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
		global.Lock1 <- true
		if global.Broadcasted[bid] {
			return
		}
		global.Broadcasted[bid] = true
		<-global.Lock1

		// 更新 Cost
		global.Lock2 <- true
		UpdateCost(source, parts[3:])
		<-global.Lock2

		// 向其他路由器继续转发
		Broadcast(msg, source)
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
	msg += "......From " + strconv.Itoa(global.Port)
	fmt.Fprintf(conn, msg)
	conn.Close()
}
