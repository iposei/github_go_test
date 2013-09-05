package main

import (
	"fmt"
	"net"
	"os"
)

var conn net.Conn

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8800")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var count int = 0

	for {
		conn, err = listener.Accept()
		fmt.Println(count)
		count++

		if err != nil || count == 5 {

			break
		}
		//go handleClient(conn)
	}
}
