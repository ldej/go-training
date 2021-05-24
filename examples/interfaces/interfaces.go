package main

import "fmt"

type Namer interface {
	Name() string
}

type Student struct {
	name string
}

func (s Student) Name() string {
	return s.name
}

type Teacher struct {
	name  string
	title string
}

func (t Teacher) Name() string {
	return t.name
}

func GetName(n Namer) string {
	return n.Name()
}

func main() {
	student := Student{name: "student"}
	teacher := Teacher{name: "teacher"}

	fmt.Println(GetName(student))
	fmt.Println(GetName(teacher))
}
