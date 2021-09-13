package main

import "fmt"

func main() {
	fmt.Println("Введите название месяца")
	var word string
	_, _ = fmt.Scan(&word)
	switch word {
	case "декабрь", "январь", "февраль":
		fmt.Println("это зима")
	case "март", "апрель", "май":
		fmt.Println("это весна")
	case "июнь", "июль", "август":
		fmt.Println("это лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("это осень")
	default:
		fmt.Println("Такого месяца я не знаю")
	}
}
