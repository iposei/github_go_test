package main

import (
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8800")
	_, _ = conn.Write([]byte("hello"))
}
