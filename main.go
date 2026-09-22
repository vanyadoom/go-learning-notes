package main

import (
	"fmt"
	"strconv"
)

func ProcessInput(input string) (string, error) {
	n, err := strconv.Atoi(input)
	if err != nil {
		return "", fmt.Errorf("Не удалось прочитать число: %w", err)
	}
	if n%2 == 0 {
		return "Чётное", nil
	}
	return "Нечётное", nil
}

func main() {
	result, err := ProcessInput("42")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", result) // Выведет: Результат: Чётное
}
