package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade int
}

func (s Student) Print() {
	fmt.Println(s.Name, s.Age, s.Grade)
}

func (s Student) IsAdult() bool {
	if s.Age >= 18 {
		return true
	}
	return false
}

func (s *Student) SetGrade(grade int) {
	s.Grade = grade
}

func main() {

	student := Student{
		Name:  "Ivan",
		Age:   25,
		Grade: 5,
	}

	student.Print()
	fmt.Println("Совешеннолетний: ", student.IsAdult())

	student.SetGrade(4)
	student.Print()

}
