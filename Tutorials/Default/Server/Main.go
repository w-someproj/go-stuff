package main

import (
	"fmt"
	"net"
)

func main() {
	message := `message from server`
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(`Not started, cause `, err)
		return
	}
	defer listener.Close()
	fmt.Println(`Server is listening`)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(`Not accept, cause `, err)
			return
		}
		conn.Write([]byte(message))
		conn.Close()
	}
}
