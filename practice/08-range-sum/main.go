package main

import "fmt"

func main() {
	var a int
	var b int
	sum := 0

	fmt.Scan(&a, &b)

	if a > b {
		a, b = b, a
	}

	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Println("Сумма чисел от", a, "до", b, ": ", sum)
}
