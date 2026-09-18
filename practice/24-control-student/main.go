package main

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

//сохраняю работу как она есть, чтобы позднее проработать ошибки
