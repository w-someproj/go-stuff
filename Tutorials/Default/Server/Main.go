package main

import (
	"fmt"
	"net"
)

var dict = map[string]string{
	"red":    "hat",
	"green":  "grass",
	"blue":   "sky",
	"yellow": "sun",
}

func main() {
	ConnectionWithClient()
}

func SendMessage() {
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

func ConnectionWithClient() {
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(`Not started, cause `, err)
		return
	}
	defer listener.Close()
	fmt.Println(`Server is listening`)
	for {
		conn, err := listener.Accept()
		if err != nil { // if error drop connection
			fmt.Println(`Not accept, cause `, err)
			conn.Close()
			continue
		}
		go HandleClientRequest(conn)
	}
}

func HandleClientRequest(conn net.Conn) {
	defer conn.Close()
	for {
		// get data from request
		input := make([]byte, 1024*4)
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		source := string(input[0:n])
		// get assoсiation
		target, ok := dict[source]
		if !ok {
			target = "undefine word to create association"
		}
		fmt.Println(source, "-", target)
		conn.Write([]byte(target))
	}
}
