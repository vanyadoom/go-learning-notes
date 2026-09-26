package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Scanln(&text)

	count := 0

	for range text {
		count++
	}
	fmt.Println(count)
}
