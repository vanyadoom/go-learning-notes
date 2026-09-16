package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade int
}

func createStudent(name string, age, grade int) Student {
	return Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
}

func main() {
	var name string
	var age, grade int
	fmt.Scan(&name, &age, &grade)

	student := createStudent(name, age, grade)

	fmt.Println(student.Name)
	fmt.Println(student.Age)
	fmt.Println(student.Grade)
}
