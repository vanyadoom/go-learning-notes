package main

import "fmt"

func main() {
	numbers := []int{4, 8, 12, 16}

	index := len(numbers) - 2

	fmt.Println(index)
	fmt.Println(numbers[index])
}
