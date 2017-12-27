package utils

import (
	"net"
)

// Listen .
func Listen(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	CheckErr(err)
	for {
		conn, err := ln.Accept()
		CheckErr(err)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
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
	}
}
