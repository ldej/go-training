package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "your-name")
	flag.Parse() // important, otherwise it doesn't work

	if name == "" {
		flag.PrintDefaults()
		log.Fatal("no name provided")
	}
	fmt.Println("Welcome:")
	fmt.Println(name)
}
