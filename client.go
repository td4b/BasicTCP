package main

import (
	"bufio"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//writer = bufio.NewWriter(conn)


	//conn.Write()

}
