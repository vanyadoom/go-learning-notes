package main

import "fmt"

func main() {
	var text string

	fmt.Scan(&text)

	for i, r := range text {
		if r == 'Я' {
			fmt.Println(i)
		}
	}
}
