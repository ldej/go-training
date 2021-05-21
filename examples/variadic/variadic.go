package main

import "fmt"

func Add(a ...int) int {
	sum := 0
	for _, b := range a {
		sum += b
	}
	return sum
}

func main() {
	fmt.Println(Add(1, 2, 3))
}
