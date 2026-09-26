package main

import "fmt"

func main() {
	var text, s, sim string

	fmt.Scan(&text, &s, &sim)

	result := ""

	for _, r := range text {
		if string(r) == s {
			result += sim
		} else {
			result += string(r)
		}
	}

	fmt.Println(result)
}
