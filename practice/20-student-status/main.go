package main

import (
	"errors"
	"fmt"
)

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

func (s Student) CanGraduate() bool {
	return s.isAdult() && s.isGradeAtLeast(4)
}

func (s *Student) SetGrade(grade int) error {
	if grade < 1 || grade > 5 {
		return errors.New("Неверная оценка")
	}
	s.Grade = grade
	return nil
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
	status := student.Status()

	oldGrade := student.Grade
	err := student.SetGrade(5)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	excellent := student.IsExcellent()
	canGraduate := student.CanGraduate()

	fmt.Println("Имя: ", student.Name)
	fmt.Println("Статус: ", status)
	fmt.Println("Оценка не ниже 4: ", gradeOK)
	fmt.Println("Старая оценка: ", oldGrade)
	fmt.Println("Новая оценка: ", student.Grade)
	fmt.Println("Отличник: ", excellent)
	fmt.Println("Можно окончить обучение: ", canGraduate)
}
