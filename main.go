package main

import "fmt"

func main() {

	x := 10

	if true {
		x := 20
		fmt.Println(x)
	}
	fmt.Println(x)
}
