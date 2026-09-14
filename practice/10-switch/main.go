package main

import "fmt"

func main() {
	var day int

	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("Среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница")
	case 6:
		fmt.Println("Суббота")
	case 7:
		fmt.Println("Воскресенье")
	default:
		fmt.Println("Ошибка: номер дня должен быть от 1 до 7.")
	}

	if day >= 1 && day <= 5 {
		fmt.Println("Рабочий день")
	} else if day >= 6 && day <= 7 {
		fmt.Println("Выходной")
	}
}
