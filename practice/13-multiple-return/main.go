package main

import "fmt"

func calculate(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	var a, b int
	fmt.Print("Введите два числа: ")
	fmt.Scan(&a, &b)
	sum, umno := calculate(a, b)
	fmt.Println("Сумма:", sum)
	fmt.Println("Произведение:", umno)

}
