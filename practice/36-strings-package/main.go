package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string = "Go go Go"
	var word string = "go"

	text = strings.ToLower(text)
	word = strings.ToLower(word)

	if strings.Contains(text, word) {
		fmt.Println(strings.Count(text, word))
	} else {
		fmt.Println(0)
	}

}
