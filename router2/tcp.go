package router2

import (
	"fmt"
	"net"
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
	case "N":
		u, _ := strconv.Atoi(parts[1])
		lock.Lock()
		updateCost(u, parts[2:])
		lock.Unlock()
		break
	case "Q":
		u, _ := strconv.Atoi(parts[1])
		v, _ := strconv.Atoi(parts[2])
		next := answer(u, v)
		msg := "A|" + strconv.Itoa(next)
		send(u, msg)
		break
	case "A":
		res, _ := strconv.Atoi(parts[1])
		if res == -1 {
			util.Prompt("错误: 找不到下一跳路由器")
			break
		}
		next <- res
		break
	case "R":
		dest, _ := strconv.Atoi(parts[2])
		if dest == port {
			break
		}
		query(port, dest)
		send(<-next, msg)
		break
	case "D":
		p, _ := strconv.Atoi(parts[1])
		delete(all, p)
		delete(near, p)
		for _, u := range near {
			delete(u, p)
		}
		delete(dist, p)
		for _, u := range dist {
			delete(u, p)
		}
		delete(prev, p)
		for _, u := range prev {
			delete(u, p)
		}
		delete(cost, p)
		for _, u := range cost {
			delete(u, p)
		}
		updated = true
		break
	default:
		break
	}
}

func send(p int, msg string) {
	conn, err := net.Dial("tcp", "0.0.0.0:"+strconv.Itoa(p))
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, msg)
	conn.Close()
	util.Log("发送: %d %s", p, msg)
}

func testPort(p int) bool {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(p))
	if err != nil {
		return false
	}
	ln.Close()
	return true
}
