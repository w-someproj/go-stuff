package main

import "fmt"

func main() {
	//Hello()
	//Kvadrat()
	//MultiplyTable()
	Sum(1, 2, 3)
	Sum(1, 2, 3, 4, 5)
}

func Hello() {
	fmt.Println(`Hello`)
	fmt.Println(`Hello World`)
}

func Kvadrat() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, ` ^ 2 = `, i*i)

	}
}

func MultiplyTable() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " | ")
		for j := 1; j <= 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}
}

func Sum(numbers ...int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
}
