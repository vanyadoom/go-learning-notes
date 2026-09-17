package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade int
}

func (s Student) isAdult() bool {
	return s.Age >= 18
}

func (s Student) isGradeAtLeast(minGrade int) bool {
	return s.Grade >= minGrade
}

func (s *Student) SetGrade(grade int) {
	s.Grade = grade
}

func (s *Student) SetAge(age int) {
	s.Age = age
}

func main() {
	var name string
	var age, grade int
	fmt.Scan(&name, &age, &grade)

	student := Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}

	adult := student.isAdult()
	gradeOK := student.isGradeAtLeast(4)
	oldGrade := student.Grade
	oldAge := student.Age

	student.SetAge(30)

	student.SetGrade(5)

	fmt.Println("Имя: ", student.Name)
	fmt.Println("Совершеннолетний: ", adult)
	fmt.Println("Оценка не ниже 4: ", gradeOK)
	fmt.Println("Старая оценка: ", oldGrade)
	fmt.Println("Новая оценка: ", student.Grade)
	fmt.Println("Старый возраст: ", oldAge)
	fmt.Println("Новый возраст: ", student.Age)
}
