package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade int
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
	gradeOK := student.isGradeAtLeast(4)

	fmt.Println("Совершеннолетний: ", adult)
	fmt.Println("Оценка не ниже 4: ", gradeOK)
}
