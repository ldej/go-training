package main

import (
	"fmt"
	"time"
)

func sendMsg(c chan string) {
	time.Sleep(100 * time.Millisecond)
	c <- "Put your helmet on"
}

func main() {
	tick := time.Tick(800 * time.Millisecond)
	boom := time.After(3 * time.Second)
	msgChannel := make(chan string)
	go sendMsg(msgChannel)

	for {
		select { // blocking until msg received on one of its channels
		case msg := <-msgChannel:
			fmt.Printf("msg: %s\n", msg) // stay in loop
		case <-tick:
			fmt.Println("tick.") // stay in loop
		case <-boom:
			fmt.Println("BOOM!")
			return // abort loop
		}
	}
}
