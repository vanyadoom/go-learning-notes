package main

import "fmt"

func main() {
	var s string

	fmt.Scan(&s)

	fmt.Println("Строка: ", s)
	fmt.Println("Длина в байтах: ", len(s))
	fmt.Println("Количество символов: ", len([]rune(s)))
	fmt.Println("Первый байт: ", s[0])
	fmt.Println("Первый символ: ", []rune(s)[0])
}
