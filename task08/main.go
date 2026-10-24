package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Println(n, "— простое число")
	} else {
		fmt.Println(n, "— не простое число")
	}
}
