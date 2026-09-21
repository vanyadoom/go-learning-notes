package main

import "fmt"

func main() {
	var name string

	fmt.Scan(&name)

	fmt.Println("Привет, ", name, "!")
	fmt.Println("Длина имени: ", len(name))
	fmt.Println("Первый байт: ", name[0])

}
