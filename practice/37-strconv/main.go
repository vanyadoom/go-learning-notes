package main

import (
	"fmt"
	"strconv"
)

func main() {

	text := "25"

	num, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println("Ошибка")
		return
	}
	result := num + 5

	fmt.Println(result)

	textResult := strconv.Itoa(result)

	fmt.Println(textResult)
}
