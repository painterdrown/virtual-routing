package router

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/painterdrown/virtual-routing/util"
)

// Listen 开始监听端口。
func Listen() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	util.CheckErr(err)
	for {
		conn, err := ln.Accept()
		util.CheckErr(err)
		go handleConn(&conn)
	}
}

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
	util.Log("接收: %s", msg)
	handleMsg(msg)
}

func handleMsg(msg string) {
	if len(msg) == 0 {
		return
	}
	parts := strings.Split(msg, "|")
	op := parts[0]

	switch op {
	case "B":
		// 判断该广播信息是否已经被该主机广播过
		bid, _ := strconv.ParseInt(parts[1], 10, 64)
		lock1.Lock()
		if broadcasted[bid] {
			lock1.Unlock()
			break
		}
		broadcasted[bid] = true
		lock1.Unlock()
		source, _ := strconv.Atoi(parts[2])
		// 更新 Cost
		lock2.Lock()
		updateCost(source, parts[3:])
		lock2.Unlock()
		// 向其他路由器继续转发
		broadcast(msg)
		break

	case "R":
		dest, _ := strconv.Atoi(parts[2])
		if dest == port {
			util.Log("接收: %s", msg)
		} else {
			forward(dest, msg)
		}
		break

	case "D":
		// 判断该广播信息是否已经被该主机广播过
		did, _ := strconv.ParseInt(parts[1], 10, 64)
		lock1.Lock()
		if broadcasted[did] {
			lock1.Unlock()
			break
		}
		broadcasted[did] = true
		lock1.Unlock()
		downport, _ := strconv.Atoi(parts[2])
		all[downport] = false
		if near[downport] {
			near[downport] = false
		}
		prev[downport] = -1
		cost[port][downport] = bigenough
		dist[downport] = bigenough
		updated = true
		// 向其他路由器继续转发
		broadcast(msg)
		break
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

func testPort(p int) bool {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(p))
	if err != nil {
		return false
	}
	ln.Close()
	return true
}

func forward(dest int, msg string) {
	var before int
	for {
		before = prev[dest]
		if before == port {
			send(dest, msg)
		} else if before == -1 {
			util.Log("转发（找不到下一跳路由器）: %s", msg)
			break
		} else {
			dest = before
		}
	}
}
