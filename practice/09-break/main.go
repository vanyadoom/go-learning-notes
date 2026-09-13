package main

import "fmt"

func main() {
	var n int

	sum := 0
	count := 0

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		count++
		sum += n
	}
	fmt.Println("Сумма: ", sum)
	fmt.Println("Количество чисел: ", count)
}
