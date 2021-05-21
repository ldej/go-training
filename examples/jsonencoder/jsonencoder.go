package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./student.json")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var student Student
	err = json.NewDecoder(f).Decode(&student)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(student)

	student.Skills = append(student.Skills, Skill{Name: "json"})

	o, err := os.Create("./output.json")
	defer o.Close()
	if err != nil {
		log.Fatal(err)
	}
	enc := json.NewEncoder(o)
	enc.SetIndent("", "  ") // pretty print
	err = enc.Encode(student)
	if err != nil {
		log.Fatal(err)
	}
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
