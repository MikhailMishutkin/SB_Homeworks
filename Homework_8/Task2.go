package main

import "fmt"

func main() {
	fmt.Println("Введите день недели")
	var day string
	_, _ = fmt.Scan(&day)
	switch day {
	case "пн":
		fmt.Println("вторник")
		day = "вт"
		fallthrough
	case "вт":
		fmt.Println("среда")
		day = "ср"
		fallthrough
	case "ср":
		fmt.Println("четверг")
		day = "чт"
		fallthrough
	case "чт":
		fmt.Println("пятница")
	case "пт":
		fmt.Println("Пятница - последний рабочий день")
	default:
		fmt.Println("Такого дня недели я не знаю")
	}
}
