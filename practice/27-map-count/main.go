package main

import "fmt"

func main() {

	students := map[string]int{

		"Ivan":  5,
		"Alex":  4,
		"Maria": 5,
		"Oleg":  3,
		"Anna":  4,
	}

	count1 := 0
	count2 := 0

	for _, value := range students {
		if value == 4 {
			count1++
		} else if value == 5 {
			count2++
		}
	}
	fmt.Println("Оценок 4: ", count1)
	fmt.Println("Оценок 5: ", count2)
}
