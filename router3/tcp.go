package router3

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/painterdrown/virtual-routing/util"
)

func listen() {
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
	case "R":
		dest, _ := strconv.Atoi(parts[2])
		if dest != port {
			n := next[dest]
			send(n, msg)
		}
		break
	case "S":
		source, _ := strconv.Atoi(parts[1])
		lock1.Lock()
		updateCost(source, parts[2:])
		lock1.Unlock()
		if updated {
			shareDist()
			updated = false
		}
		break
	case "D":
		tellNeighbors(msg)
		os.Exit(0)
		break
	default:
		break
	}
}

func send(p int, msg string) {
	conn, err := net.Dial("tcp", "0.0.0.0:"+strconv.Itoa(p))
	if err != nil {
		util.Log("错误: %s", err.Error())
		return
	}
	fmt.Fprintf(conn, msg)
	util.Log("发送: %d %s", p, msg)
	conn.Close()
}

func testListen(p int) bool {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(p))
	if err != nil {
		return false
	}
	ln.Close()
	return true
}

func testConnection(p int) bool {
	conn, err := net.Dial("tcp", "0.0.0.0:"+strconv.Itoa(p))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
