package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade int
}

func (s *Student) SetGrade(grade int) {
	s.Grade = grade
}

func (s *Student) SetAge(age int) {
	s.Age = age
}

func (s Student) isGradeAtLeast(minGrade int) bool {
	return s.Grade >= minGrade
}

func (s Student) isAdult() bool {
	return s.Age >= 18
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

	adult := student.isAdult()
	oldGrade := student.Grade
	oldAge := student.Age

	student.SetAge(30)
	student.SetGrade(5)

	fmt.Println("Совершеннолетний: ", adult)
	fmt.Println("Старая оценка: ", oldGrade)
	fmt.Println("Новая оценка: ", student.Grade)
	fmt.Println("Старый возраст: ", oldAge)
	fmt.Println("Новый возраст: ", student.Age)
}
