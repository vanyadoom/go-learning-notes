package main

import "fmt"

func main() {
	numbers := []int{}

	numbers = append(numbers, 10)
	numbers = append(numbers, 20)
	numbers = append(numbers, 30)

	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	numbers = append(numbers, 40)
	numbers = append(numbers, 50)

	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}
