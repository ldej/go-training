package main

import "fmt"

type School struct {
	Name     string
	Teachers []Teacher
	Students []Student
	Rooms    []Room
}

type Teacher struct {
	Title string
	Name  string
}

type Student struct {
	Name string
}

type Room struct {
	Number int
}

type Class struct {
	Name     string
	Teacher  Teacher
	Students []Student
	Room     Room
}

func main() {
	school := School{
		Name: "Go school",
	}

	laurence := Teacher{
		Name:  "Laurence",
		Title: "Engineer",
	}
	akshay := Teacher{
		Name:  "Akshay",
		Title: "Master",
	}

	teachers := []Teacher{laurence, akshay}
	school.Teachers = teachers

	student1 := Student{
		Name: "Student 1",
	}
	student2 := Student{
		Name: "Student 2",
	}
	students := []Student{student1, student2}
	school.Students = students

	school.Rooms = []Room{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
	}

	class := Class{
		Name:     "Science",
		Teacher:  laurence,
		Students: students,
		Room:     school.Rooms[0],
	}
	fmt.Println(class)
}
