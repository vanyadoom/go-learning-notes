package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	var a, b int

	fmt.Print("Введите два числа: ")
	fmt.Scan(&a, &b)
	
	fmt.Println("Сумма двух чисел: ", add(a, b))
}
