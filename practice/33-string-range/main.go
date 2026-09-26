package main

import (
	"fmt"
)

func main() {
	var r string
	fmt.Scan(&r)

	count := 0

	for _, r := range r {
		if r == 'а' {
			count++
		}
	}
	fmt.Println(count)
}
