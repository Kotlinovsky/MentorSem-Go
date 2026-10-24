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

	sum, even, odd := 0, 0, 0
	var max, min int

	for i := 1; i <= n; i++ {
		var x int
		fmt.Printf("Число %d: ", i)
		fmt.Scan(&x)

		sum += x

		if i == 1 || x > max {
			max = x
		}
		if i == 1 || x < min {
			min = x
		}

		if x%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	average := float64(sum) / float64(n)

	fmt.Println("Сумма:", sum)
	fmt.Printf("Среднее: %f\n", average)
	fmt.Println("Максимум:", max)
	fmt.Println("Минимум:", min)
	fmt.Println("Чётных:", even)
	fmt.Println("Нечётных:", odd)
}
