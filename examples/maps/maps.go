package main

import (
	"fmt"
)

func main() {
	studentsOnSchool := map[string][]string{
		"Cambridge": {"Raj", "Alice"},
		"MIT":       {"Bob"},
	}
	fmt.Printf("1: %+v\n", studentsOnSchool) // %+v debugging convenience

	studentsOnSchool["DU"] = []string{"Abhi"} // add map entry
	fmt.Printf("2: %+v\n", studentsOnSchool)  // %+v debugging convenience

	delete(studentsOnSchool, "MIT")          // remove map entry
	fmt.Printf("3: %+v\n", studentsOnSchool) // %+v debugging convenience

	cambridgeStudents, found := studentsOnSchool["Cambridge"] // get map entry
	if !found {
		cambridgeStudents = []string{}
	}
	cambridgeStudents = append(cambridgeStudents, "Neha")
	studentsOnSchool["Cambridge"] = cambridgeStudents // put map entry

	for key, value := range studentsOnSchool { // iterate map
		fmt.Printf("4: %s - %v\n", key, value)
	}
}
