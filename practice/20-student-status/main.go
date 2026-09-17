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

func (s Student) IsExcellent() bool {
	return s.Grade >= 5
}

func (s *Student) SetGrade(grade int) {
	s.Grade = grade
}

func (s *Student) SetAge(age int) {
	s.Age = age
}

func (s Student) Status() string {
	if s.isAdult() {
		return "Совершеннолетний"
	}
	return "Не совершеннолетний"
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
	gradeOK := student.isGradeAtLeast(4)
	excellent := student.IsExcellent()
	oldGrade := student.Grade

	status := student.Status()

	student.SetGrade(5)

	fmt.Println("Имя: ", student.Name)
	fmt.Println("Статус: ", status)
	fmt.Println("Оценка не ниже 4: ", gradeOK)
	fmt.Println("Старая оценка: ", oldGrade)
	fmt.Println("Новая оценка: ", student.Grade)
	fmt.Println("Отличник: ", excellent)
}
