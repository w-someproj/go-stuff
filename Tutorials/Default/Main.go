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

	//f := Square()
	//fmt.Println(f())
	//fmt.Println(f())

	//fmt.Println(Factorial(4))
	//fmt.Println(Fibonacci(4))

	//defer Finish()
	//fmt.Println(`Start`)
	//fmt.Println(Divide(15, 0))

	//Slice()
	//Map()

	//Pointers()

	Structs()
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

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Fibonacci(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return Fibonacci(x-1) + Fibonacci(x-2)
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

func Slice() {
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users := initialUsers[2:7]
	fmt.Println(users)
	users = append(users[:1], users[3:]...)
	fmt.Println(users)
}

func Map() {
	var people = map[string]int{
		"Bob":   1,
		"Alice": 2,
		"Kate":  3,
		"Sam":   4,
	}
	fmt.Println(people["Alice"])
	fmt.Println(people["Bob"])
	people["Bob"] = 5
	fmt.Println(people["Bob"])

	if val, ok := people["Kate"]; ok {
		fmt.Println(val)
	}

	for key, value := range people {
		fmt.Println(key, value)
	}
	delete(people, "Bob")
	fmt.Println(people)
}

func Pointers() {
	var x int = 4
	var p *int
	p = &x
	fmt.Println(p)
	fmt.Println(*p)
	*p = 10
	fmt.Println(x)

	x = 4
	fmt.Println("x before:", x)
	ChangeValue(&x)
	fmt.Println("x after:", x)
	var people = map[string]int{
		"Bob":   1,
		"Alice": 2,
		"Kate":  3,
		"Sam":   4,
	}
	fmt.Println(people)
	ChangeMap(people)
	fmt.Println(people)
}

func ChangeValue(x *int) {
	*x = 15
}

func ChangeMap(values map[string]int) {
	if val, ok := values[`Bob`]; ok {
		val = 10
		fmt.Println(val)
		values[`Bob`] = 10
	}
	delete(values, `Sam`)
}

func Structs() {
	tom := Person{Name: `Tom`, Age: 24, Contact: ContactInfo{}}
	bob := Person{Name: `Bob`, Age: 24}
	fmt.Println(tom.Name)
	fmt.Println(bob)
	tomPointer := &tom
	tomPointer.Age = 29
	fmt.Println(tom.Age)
	(*tomPointer).Age = 32
	fmt.Println(tom.Age)
	tom.Email = `tom`
	fmt.Println(tom.Email)
	tom.Contact.Email = `supertom`
	fmt.Println(tom.Email)
	tom.Print()
	tom.ChangeAge(17)
	tom.Print()
}

type Person struct {
	Name string
	Age  int
	ContactInfo
	Contact ContactInfo
}

type ContactInfo struct {
	Email string
	Phone string
}

func (p Person) Print() {
	fmt.Println(p.Name, p.Age)
}

func (p *Person) ChangeAge(newAge int) {
	p.Age = newAge
}
