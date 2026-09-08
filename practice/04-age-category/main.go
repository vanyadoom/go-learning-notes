package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Scan(&name, &age)

	fmt.Println("Привет", name, "!")

	if age < 13 {
		fmt.Println("Категория: Ребёнок")
	} else if age < 18 {
		fmt.Println("Категория: Подросток")
	} else if age < 65 {
		fmt.Println("Категория: Взрослый")
	} else {
		fmt.Println("Категория: Пенсионер")
	}
}
