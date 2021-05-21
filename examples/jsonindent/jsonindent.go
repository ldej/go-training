package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	a := Student{
		Name:   "Laurence",
		Age:    31,
		Family: []string{"father", "mother"},
		Skills: []Skill{
			{Name: "Programming"},
			{Name: "Teaching"},
		},
	}

	studentJSON, err := json.MarshalIndent(a, "", "  ") // indent makes it pretty
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(studentJSON))
}

type Student struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Family []string `json:"family"`
	Skills []Skill  `json:"skills"`
}

type Skill struct {
	Name string `json:"name"`
}
