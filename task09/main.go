package main

import "fmt"

func main() {
	var balance float64
	fmt.Print("Введите баланс: ")
	fmt.Scan(&balance)

	for {
		fmt.Println("\n1 — Посмотреть баланс")
		fmt.Println("2 — Пополнить баланс")
		fmt.Println("3 — Снять средства")
		fmt.Println("0 — Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Баланс:", balance)
		case 2:
			var amount float64
			fmt.Print("Сумма пополнения: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Сумма должна быть больше нуля")
			} else {
				balance += amount
				fmt.Println("Баланс пополнен. Текущий баланс:", balance)
			}
		case 3:
			var amount float64
			fmt.Print("Сумма снятия: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Сумма должна быть больше нуля")
			} else if amount > balance {
				fmt.Println("Недостаточно средств")
			} else {
				balance -= amount
				fmt.Println("Выдано:", amount, "Остаток:", balance)
			}
		case 0:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неизвестное действие")
		}
	}
}
