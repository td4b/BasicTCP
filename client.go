package main

import (
	"bufio"
	"net"
)


// Define the Protocol Handshake.
var (
	CH = []byte("0x001")
	SH = []byte("0x002")
	CA = []byte("0x003")
	SA = []byte("0x004")
)

func write(input string,conn net.Conn) {
	bufio.NewWriter(conn)
	conn.Write([]byte(input))
}

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	write("Meow",conn)
}
