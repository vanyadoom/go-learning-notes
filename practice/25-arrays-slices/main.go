package main

import "fmt"

func main() {
	numbers := []int{}

	numbers = append(numbers, 10)
	numbers = append(numbers, 20)
	numbers = append(numbers, 30)
	numbers = append(numbers, 40)
	numbers = append(numbers, 50)

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	fmt.Println("Количество элементов:", len(numbers))

	numbers[2] = 300

	for i := 0; i < len(numbers); i++ {
		fmt.Println("После изменения: ", numbers[i])
	}
}
