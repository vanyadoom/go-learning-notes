package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Scanln(&text)

	count := 0

	for _, r := range text {
		fmt.Println(string(r))
		count++
	}
	fmt.Println("Количество символов: ", count)
}
