package main

import (
	"time"
)

var number int

func Increase(id string) {
	for true {
		number++
		//fmt.Println(id, number)
	}
}

func run() {
	go Increase("1")
	go Increase("2")
	<-time.Tick(1 * time.Millisecond)
}

func main() {
	// Run this with go run -race to see:
	// Found 1 data race(s)
	run()
}
