package main

import (
	"fmt"
)

// functions are first-class citizens that can be passed around like variables and objects
func main() {
	fmt.Println(Do(1, 2, Add))
	fmt.Println(Do(2, 1, Sub))

	c := 10

	// define function in-line and put in variable
	mul := func(a int, b int) int {
		return a * b * c
	}
	fmt.Println(mul(2, 2))

	fmt.Println(Do(3, 3, mul))

	// anonymous function, also called a lambda function, closure (function)
	fmt.Println(Do(10, 10, func(a int, b int) int {
		return a / b
	}))

	go Add(1, 2)
}

func Do(a int, b int, f func(c int, d int) int) int {
	return f(a, b)
}

func Add(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}
