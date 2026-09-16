package main

import (
	"errors"
	"fmt"
)

func checkNumber(n int) (int, error) {
	if n > 0 {
		return n * 2, nil
	}
	return 0, errors.New("Число должно быть положительным")
}

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	result, err := checkNumber(n)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Результат:", result)
}
