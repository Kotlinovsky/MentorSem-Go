package main

import "fmt"

func main() {
	var a, b float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Print("Введите операцию (+, -, *, /): ")
	fmt.Scan(&op)

	switch op {
	case "+":
		fmt.Println("Результат:", a+b)
	case "-":
		fmt.Println("Результат:", a-b)
	case "*":
		fmt.Println("Результат:", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно")
		} else {
			fmt.Println("Результат:", a/b)
		}
	default:
		fmt.Println("Неизвестная операция")
	}
}
