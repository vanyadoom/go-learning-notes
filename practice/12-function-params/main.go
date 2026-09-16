package main

import "fmt"

func square(a int) int {
	return a * a
}

func main() {
	var a int
	fmt.Print("Введите число: ")
	fmt.Scan(&a)
	result := square(a)
	fmt.Println("Квадрат:", result)
}
