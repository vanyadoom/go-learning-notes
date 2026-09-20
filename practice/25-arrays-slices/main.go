package main

import "fmt"

func main() {

	numbers := make([]int, 3, 5)

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	numbers = append(numbers, 40)
	numbers = append(numbers, 50)

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(len(numbers))
}
