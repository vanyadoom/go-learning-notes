package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	if a > b {
		a, b = b, a
	}

	sum := 0
	count := 0

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			sum += i
			count++
		}
	}

	fmt.Println("Сумма чётных:", sum)
	fmt.Println("Количество чётных:", count)
}
