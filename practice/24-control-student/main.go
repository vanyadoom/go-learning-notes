package main

import (
	"fmt"
)

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

func isAdult(student Student) bool {
	if student.Age >= 18 {
		return true
	}
	return false
}

func SetGrade(student *Student, grade int) {
	if grade < 1 || grade > 5 {
		return fmt.Errorf("Неверная оценка")
	}
	student.Grade = grade
	return nil
}

func Status(student Student) string {
	if student.Grade == 5 {
		return "Отличник"
	} else if student.Grade == 4 {
		return "Хорошист"
	}
	return "Обычный студент"
}

func main() {
	var name string
	var age int
	var grade int

	fmt.Scan(&name, &age, &grade)
	student := createStudent(name, age, grade)

	setGrade := SetGrade(student, grade)
	if setGrade != nil {
		fmt.Println("Ошибка:", setGrade)
		return
	}
	fmt.Println(setGrade)
}
//сохраняю работу как она есть, чтобы позднее проработать ошибки

