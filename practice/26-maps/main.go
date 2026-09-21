package main

import "fmt"

func main() {

	students := map[string]int{
		"Ivan":  5,
		"Alex":  4,
		"Maria": 5,
	}

	var name string

	fmt.Scan(&name)

	value, ok := students[name]

	if ok {
		fmt.Println("Оценка: ", value)
	} else {
		fmt.Println("Студент не найден")
	}
}
