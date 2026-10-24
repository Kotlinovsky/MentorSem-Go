package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите N: ")
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("N должно быть больше 0")
		return
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	average := float64(sum) / float64(n)
	fmt.Println("Сумма:", sum)
	fmt.Printf("Среднее: %f\n", average)
}
