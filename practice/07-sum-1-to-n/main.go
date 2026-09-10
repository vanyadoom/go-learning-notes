package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	evenSum := 0
	oddSum := 0
	for i := 1; i <= n; i++ {
		sum += i
		if i%2 == 0 {
			evenSum = evenSum + i
		} else {
			oddSum = oddSum + i
		}
	}
	fmt.Println("Сумма чисел от 1 до", n, ":", sum)
	fmt.Println("Сумма чётных чисел: ", evenSum)
	fmt.Println("Сумма нечётных чисел: ", oddSum)
}
