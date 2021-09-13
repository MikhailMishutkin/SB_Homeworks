package main

import "fmt"

func main() {
	fmt.Println("Введите сумму дохода")
	var income, tax float64
	_, _ = fmt.Scan(&income)
	switch {
	case income >= 70000:
		tax = (income - 70000) * 50 / 100
		fmt.Println("Сумма налога свыше 70000 составит 50% от дохода и равна ", tax)
		income = 70000
		fallthrough
	case income >= 50000:
		tax = (income - 50000) * 30 / 100
		fmt.Println("Сумма налога от 50000 до 70000 составит 30% от дохода и равна ", tax)
		income = 50000
		fallthrough
	case income >= 10000:
		tax = (income - 10000) * 20 / 100
		fmt.Println("Сумма налога от 10000 до 50000 составит 20% от дохода и равна ", tax)
		income = 10000
		fallthrough
	case income > 0:
		tax = income * 13 / 100
		fmt.Println("Сумма налога до 10000 составит 13% от дохода и равна ", tax)
	default:
		fmt.Println("Доход должен быть положительным")
	}

}
