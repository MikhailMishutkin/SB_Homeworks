package main

import (
	"fmt"
)

func main() {
	fmt.Println("Шахматная доска")
	fmt.Println("===================")
	var a int
	fmt.Println("Введите количество клеток доски")
	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		var j int
		j = i % 2
		for j < a {
			if j%2 > 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
			j++
		}
		fmt.Println()
	}
}
