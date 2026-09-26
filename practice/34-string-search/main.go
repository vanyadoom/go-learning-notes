package main

import (
	"fmt"
)

func main() {
	var text string

	var s string

	fmt.Scan(&text)
	fmt.Scan(&s)

	for i, r := range text {
		if r == []rune(s)[0] {
			fmt.Println(i)
			fmt.Println("нашёл")
		}
	}
}
