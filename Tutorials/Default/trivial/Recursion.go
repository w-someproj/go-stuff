package trivial

import "fmt"

func Recursion() {
	fmt.Println(Factorial(4))
	fmt.Println(Fibonacci(4))
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
