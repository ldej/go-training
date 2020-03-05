package main

import (
	"fmt"
)

//go:generate enumer -type Color -text -json -transform snake
type Color int

const (
	Unknown    Color = iota // 0 (=default)
	Red                     // 1
	Green                   // 2
	Blue                    // 3
)

func main() {
	var myColor Color // uses default
	otherColor := Green
	fmt.Printf("my-color: %v (%d), other-color: %v (%d)\n",
		myColor, myColor, otherColor, otherColor)
}
