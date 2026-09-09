package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	even := 0
	odd := 0

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i, "- чётное")
			even++
		} else {
			fmt.Println(i, "- нечётное")
			odd++
		}
	}
	fmt.Println("Количество чётных чисел:", even)
	fmt.Println("Количество нечётных чисел:", odd)
}
