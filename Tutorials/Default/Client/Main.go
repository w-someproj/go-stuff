package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	ConnectionWithSever()
}

func GetMessage() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	io.Copy(os.Stdout, conn)
	fmt.Println("\nDone")
}

func ConnectionWithSever() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		var source string
		fmt.Print(`Enter word:`)
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println(`Need word`, err)
			continue
		}
		n, err := conn.Write([]byte(source))
		if n == 0 || err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print(`Association:`)
		conn.SetReadDeadline(time.Now().Add(time.Second * 3))
		for {
			buff := make([]byte, 8)
			nRead, errRead := conn.Read(buff)
			if errRead != nil {
				break
			}
			fmt.Print(string(buff[0:nRead]))
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 500))
		}

		fmt.Println()

	}
}
