package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string
	var word string

	fmt.Scan(&text, &word)

	text = strings.ToLower(text)
	word = strings.ToLower(word)

	fmt.Println(strings.Contains(text, word))

}
