package trivial

import (
	"fmt"
	"io"
	"os"
)

func Flow() {
	//FileCreate()
	FileRead()
}

func FileCreate() {
	file, err := os.Create(`hello.txt`)
	text := `Hello, bro`
	if err != nil {
		fmt.Println(`something wrong with file: `, err)
		os.Exit(1) // exit from here
	}
	defer file.Close()
	file.WriteString(text)
	fmt.Println(file.Name())
}

func FileRead() {
	file, err := os.Open(`hello.txt`)
	if err != nil {
		fmt.Println(`something wrong with file: `, err)
		os.Exit(1) // exit from here
	}
	defer file.Close()
	data := make([]byte, 8)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
	}
}
