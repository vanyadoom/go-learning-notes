package main

import "fmt"

func main() {
	text := "ЯAБ" // Я и Б — кириллические, A — латинская

	for i, r := range text {
		fmt.Println(i, r)
	}
}


// i = 0 2 3 
// r = Я А Б 
