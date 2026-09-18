package main

import "fmt"

type Student struct {
	Age   int
	Grade int
}

func (s Student) CanGraduate() bool {
	return s.Age >= 18 && s.Grade >= 4
}

func main() {
	student := Student{
		Age:   18,
		Grade: 3,
	}

	canGraduate := student.CanGraduate()

	student.Grade = 5

	fmt.Println(canGraduate)
}
