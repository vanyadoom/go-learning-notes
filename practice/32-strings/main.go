package main

import "fmt"

func main() {
	var a string

	fmt.Scan(&a)

	runes := []rune(a)

	fmt.Println("Длина строки в байтах: ", len(a))
	fmt.Println("Количество символов: ", len([]rune(a)))

	first := runes[0]

	fmt.Println("Первый символ: ", first)

	last := runes[len(runes)-1]

	fmt.Println("Последний символ: ", last)

	var reversed string

	for i := len(runes) - 1; i >= 0; i-- {
		reversed += string(runes[i])
	}
	fmt.Println("Перевёрнутая строка: ", reversed)
}
