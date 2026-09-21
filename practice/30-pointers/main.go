package main

import "fmt"

func main() {

	x := 10

	p := &x

	fmt.Println("Значение x: ", x)
	fmt.Println("Адрес x: ", &x)
	fmt.Println("Значение через указатель: ", *p)

	*p = 20

	fmt.Println("Новое значение x: ", x)
}
