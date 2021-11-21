package trivial

import "fmt"

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
