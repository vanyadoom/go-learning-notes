package main

import "fmt"

func calculate(a, b int, operation string) (int, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Ошибка! На 0 делить нельзя!")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Неизвестная операция!")
	}
}

func main() {
	var a, b int
	var operation string
	fmt.Scan(&a, &b, &operation)
	result, err := calculate(a, b, operation)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Результат:", result)
}
