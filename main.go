package main

import "fmt"

func main() {
	var age int

	_, err := fmt.Scan(&age)

	if err != nil {
		fmt.Println("Ошибка")
		return
	}

	fmt.Println("Возраст:", age)
}
