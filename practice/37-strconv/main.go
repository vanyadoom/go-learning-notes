package main

import (
	"fmt"
	"strconv"
)

func main() {
	text := "15"
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(err)
	}
	result := num + 5
	textResult := strconv.Itoa(result)
	fmt.Println(result)
	fmt.Println(textResult)
}
