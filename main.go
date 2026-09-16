package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	sum := 0

	for {
		if n == 0 {
			break
		}

		if n%2 == 0 {
			sum += n
		}

		fmt.Scan(&n)
	}

	fmt.Println("Сумма чётных:", sum)
}
