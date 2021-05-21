package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// encoding, serialization, marshalling -> all the same -> go to json
	// decoding, deserialization, unmarshalling -> the opposite -> json to go

	a := Student{
		Name:     "Laurence",
		Age:      31,
		lastName: "de Jong",
		Family:   []string{"father", "mother"},
		Skills: []Skill{
			{Name: "Programming"},
			{Name: "Teaching"},
		},
	}

	studentJSON, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(studentJSON))

	input := `{"anything":"Laurence","family":["father","mother"],"skills":[{"name":"Programming"},{"name":"Teaching"}]}`

	var b Student
	err = json.Unmarshal([]byte(input), &b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b)
	fmt.Println(b.Name)
}

type Student struct {
	Name     string   `json:"anything"`  // this is called a struct tag
	Age      int      `json:"-"`         // ignore
	lastName string   `json:"last_name"` // this doesn't work because json.Marshal only marshals Public properties
	Family   []string `json:"family"`
	Skills   []Skill  `json:"skills"`
}

type Skill struct {
	Name string `json:"name"`
}
