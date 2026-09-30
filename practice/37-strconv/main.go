package main

import (
	"fmt"
	"strconv"
)

func main() {

	var text string

	fmt.Scanln(&text)

	num, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println("Ошибка преобразования")
		return
	} else {
		fmt.Println(num + 10)
	}
}
