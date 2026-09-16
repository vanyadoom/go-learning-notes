package main

import "fmt"

func sumEven(a, b int) int {
	sum := 0

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			sum += i
		}
	}

	return sum
}

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	result := sumEven(a, b)

	fmt.Println("Сумма чётных чисел:", result)
}
