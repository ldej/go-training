package main

import (
	"fmt"
)

func main() {
	doit()
}

func sum(a []int, resultChannel chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	resultChannel <- sum // send result back over channel
}

func doit() {
	responseChannel := make(chan int) // construct channel
	defer close(responseChannel)      // prevent resource leak

	go sum([]int{1, 2, 3}, responseChannel)      // 1 + 2 + 3 = 6
	go sum([]int{4, 5, 6}, responseChannel)      // 4 + 5 + 6 = 15
	x, y := <-responseChannel, <-responseChannel // receive from channel

	fmt.Printf("one=%d\nanother=%d", x, y) // order undefined
}
