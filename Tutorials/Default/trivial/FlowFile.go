package trivial

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Flow() {
	FileCreate()
	FileRead()
	//Copy()
	//FormatReadWrite()
	//BufioRead()
}

func FileCreate() {
	file, err := os.Create(`hello.txt`)
	defer file.Close()
	text := `Hello, bro`
	if err != nil {
		fmt.Println(`something wrong with file: `, err)
		os.Exit(1) // exit from here
	}
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

func Copy() {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	io.Copy(os.Stdout, file)
}
func FormatReadWrite() {
	filename := `people.dat`
	WriteData(filename)
	ReadData(filename)
}

func WriteData(fileName string) {
	people := []Person{
		{"Tom", 24, 60.5, ContactInfo{Email: `tom`, Phone: `123`}},
		{"Bob", 25, 60.238, ContactInfo{Email: `bob`, Phone: `456`}},
		{"Sam", 27, 70.6, ContactInfo{Email: `sam`, Phone: `789`}},
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(`something wrong with file: `, err)
		os.Exit(1) // exit from here
	}
	defer file.Close()

	for _, p := range people {
		fmt.Fprintf(file, `%s %d %.2f\n`, p.Name, p.Age, p.Weight)
	}
}

func ReadData(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(`something wrong with file: `, err)
		os.Exit(1) // exit from here
	}
	defer file.Close()
	var name string
	var age int
	var weight float64
	for {
		_, err := fmt.Fscanf(file, `%s %d %f\n`, &name, &age, &weight)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
	}
}

func BufioWrite() {
	rows := []string{
		"Smth",
		"Check",
	}

	file, err := os.Create("some.dat")
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, row := range rows {
		writer.WriteString(row)
		writer.WriteString("\n")
	}
	writer.Flush()
}

func BufioRead() {
	file, err := os.Open("some.dat")
	if err != nil {
		fmt.Println("something wrong with file: ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		fmt.Print(line)
	}
}
