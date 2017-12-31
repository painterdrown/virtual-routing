package router

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/painterdrown/virtual-routing/util"
)

// Listen 开始进行监听。
func Listen() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	util.CheckErr(err)
	for {
		conn, err := ln.Accept()
		util.CheckErr(err)
		go handleConn(&conn)
	}
}

// ======================

func handleConn(connAddr *net.Conn) {
	conn := *connAddr
	defer conn.Close()
	var buffer [512]byte
	n, err := conn.Read(buffer[0:])
	if err != nil {
		if err.Error() == "EOF" {
			return
		}
		panic(err)
	}
	msg := string(buffer[0:n])
	util.Prompt("Recieving: " + msg)
	handleMsg(msg)
}

func handleMsg(msg string) {
	if len(msg) == 0 {
		return
	}
	parts := strings.Split(msg, "|")

	// 如果是广播信息
	if parts[0] == "B" {
		bid, _ := strconv.ParseInt(parts[1], 10, 64)
		source, _ := strconv.Atoi(parts[2])

		// 判断该广播信息是否已经被该主机广播过
		lock1.Lock()
		if broadcasted[bid] {
			lock1.Unlock()
			return
		}
		broadcasted[bid] = true
		lock1.Unlock()

		// 向其他路由器继续转发
		broadcast(msg, source)

		// 更新 Cost
		lock2.Lock()
		updateCost(source, parts[3:])
		lock2.Unlock()
	}

	// 如果是路由信息
	if parts[0] == "R" {

	}
}

func send(port int, msg string) {
	conn, err := net.Dial("tcp", "0.0.0.0:"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, msg)
	conn.Close()
}
