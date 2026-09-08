package main

import "fmt"

func main() {
	var name string
	var age int

	_, err := fmt.Scan(&name, &age)

	if err != nil {
		fmt.Println("Ошибка: пожалуйста, введите корректные данные!")
		return
	}

	fmt.Println("Привет,", name, "!")

	if age >= 18 {
		fmt.Println("Ты совершеннолетний!")
	} else {
		fmt.Println("Ты несовершеннолетний!")
	}
}
