package main

import "fmt"

func main() {
	var a int
	var b int

	_, err := fmt.Scan(&a, &b)

	if err != nil {
		fmt.Println("Ошибка: пожалуйста, введите целое число!")
		return
	}

	sum := a + b

	fmt.Println("Сумма двух чисел:", sum)

}
