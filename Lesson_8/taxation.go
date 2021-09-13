package main

import "fmt"

func main() {
	fmt.Println("Введите сумму дохода")
	var income, tax float64
	fmt.Scan(&income)
	switch {
	case income >= 50000:
		tax = income * 30 / 100
		fmt.Println("Налог составит 30% от дохода и равен ", tax)
	case income >= 10000:
		tax = income * 20 / 100
		fmt.Println("Налог составит 20% от дохода и равен ", tax)
	case income > 0:
		tax = income * 13 / 100
		fmt.Println("Налог составит 13% от дохода и равен ", tax)
	case income == 0:
		fmt.Println("Да ладно 0_о")
	default:
		fmt.Println("Не может быть.....")
	}

}
