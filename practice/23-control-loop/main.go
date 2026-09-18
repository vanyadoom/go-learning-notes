package main

import "fmt"

func main() {
	var n int
	positiveCount := 0
	negativeSum := 0

	for {
		_, err := fmt.Scan(&n)

		if err != nil {
			fmt.Println("Ошибка: введите целое число")
			return
		}

		if n == 0 {
			break
		}

		if n > 0 {
			positiveCount++
		} else {
			negativeSum += n
		}
	}

	fmt.Println("Положительных чисел:", positiveCount)
	fmt.Println("Сумма отрицательных:", negativeSum)
}
