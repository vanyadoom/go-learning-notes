package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade int
}

func main() {

	student1 := Student{
		Name:  "Ivan",
		Age:   25,
		Grade: 5,
	}

	student2 := Student{
		Name:  "Alex",
		Age:   22,
		Grade: 4,
	}
	student1.Grade = 4
	fmt.Println(student1.Name, student1.Age, student1.Grade)
	fmt.Println(student2.Name, student2.Age, student2.Grade)
}
