package main

import (
	"fmt"
)

func main() {
	StartTutorial()
}

//don`t clutter main)
func StartTutorial() {
	//Hello()
	//Kvadrat()
	//MultiplyTable()
	//SumNums(1, 2, 3)
	//SumNums(1, 2, 3, 4, 5)
	//SumNums([]int{1, 2, 3}...)

	//fullName, address := GetFullInfo(`James`, `Brown`, `Green`, 1)
	//fmt.Println(fullName, "-", address)

	//var f func(int, int) int = add
	//fmt.Println(f(1,2))
	//Action(1, 2, Add)
	//Action(2, 3, Multiply)

	//slice := []int {1,2,-2,-3,4,5,-10, -5, 8}
	//sumEven := SumCriteria(slice, IsEven)
	//fmt.Println(sumEven)
	//sumPositiv := SumCriteria(slice, IsPositiv)
	//fmt.Println(sumPositiv)

	f := Square()
	fmt.Println(f())
	fmt.Println(f())
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

func SumNums(numbers ...int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
}

func GetFullInfo(firstName, lastName, street string, house int) (fullName, address string) {
	fullName = firstName + " " + lastName
	//address = street + " " + strconv.Itoa(house)
	address = fmt.Sprintf(`st.%v, h.%v`, street, house)
	return
}

func Add(x, y int) int {
	return x + y
}

func Multiply(x, y int) int {
	return x * y
}

func Action(x, y int, operation func(int, int) int) {
	result := operation(x, y)
	fmt.Println(result)
}

func IsEven(x int) bool {
	return x%2 == 0
}

func IsPositiv(x int) bool {
	return x > 0
}

func SumCriteria(slice []int, criteria func(int) bool) (sum int) {
	for _, value := range slice {
		if criteria(value) {
			sum += value
		}
	}
	return
}

func Square() func() int {
	var x = 2
	return func() int {
		x++
		return x * x
	}
}
