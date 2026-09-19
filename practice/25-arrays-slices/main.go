package main

import "fmt"

func main() {
	numbers := [5]int{}
	numbers[0] = 10
	numbers[3] = 40
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println("Количество элементов: ", len(numbers))
}
