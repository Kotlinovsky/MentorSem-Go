package main

import "fmt"

func main() {
	var n int
	fmt.Print("Сколько чисел? ")
	fmt.Scan(&n)

	positive, negative, zeros := 0, 0, 0
	var x float64

	for i := 1; i <= n; i++ {
		fmt.Printf("Число %d: ", i)
		fmt.Scan(&x)

		if x > 0 {
			positive++
		} else if x < 0 {
			negative++
		} else {
			zeros++
		}
	}

	fmt.Println("Положительных:", positive)
	fmt.Println("Отрицательных:", negative)
	fmt.Println("Нулей:", zeros)
}
