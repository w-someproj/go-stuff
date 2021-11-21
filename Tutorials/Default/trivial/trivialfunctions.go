package trivial

import "fmt"

func Trivial() {
	defer Finish()
	fmt.Println(`Start`)
	Hello()
	Kvadrat()
	MultiplyTable()
	//fmt.Println(Divide(15, 0))
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

func Finish() {
	defer fmt.Println(`Finish end`)
	fmt.Println(`Finish start`)
}

func Divide(x, y int) int {
	if y == 0 {
		panic(`Division by zero`)
	}
	return x / y
}
