package main

import "fmt"

func main() {
	//каждый охотник желает знать где сидит фазан
	// кварк окружает жаркий занавес глюонов создающих флюиды
	fmt.Println("Введите слово из считалочки про радугу")
	var word string
	_, _ = fmt.Scan(&word)
	switch word {
	case "каждый", "кварк":
		fmt.Println("красный")
	case "охотник", "окружает":
		fmt.Println("оранжевый")
	case "желает", "жаркий":
		fmt.Println("желтый")
	case "знать", "занавес":
		fmt.Println("зеленый")
	case "где", "глюонов":
		fmt.Println("голубой")
	case "сидит", "создающих":
		fmt.Println("синий")
	case "фазан", "флюиды":
		fmt.Println("фиолетовый")
	default:
		fmt.Println("Такого цвета я не знаю")
	}
}
