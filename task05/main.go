package main

import "fmt"

func main() {
	var n int
	fmt.Print("Сколько чисел? ")
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("Нужно ввести хотя бы одно число")
		return
	}

	var x float64
	fmt.Print("Число 1: ")
	fmt.Scan(&x)
	max, min := x, x

	for i := 2; i <= n; i++ {
		fmt.Printf("Число %d: ", i)
		fmt.Scan(&x)
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	fmt.Println("Максимум:", max)
	fmt.Println("Минимум:", min)
}
