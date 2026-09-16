package main

import "fmt"

func minMax(a, b int) (min int, max int) {
	if a > b {
		return b, a
	}
	return a, b
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	a, b = minMax(a, b)
	fmt.Println("Минимум:", a)
	fmt.Println("Максимум:", b)
}
