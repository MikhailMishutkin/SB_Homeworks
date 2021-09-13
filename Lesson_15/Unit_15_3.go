//Инициализация массивов

package main

import "fmt"`

func main() {
	var rain [10]int
	for i := 0; i < 10; i++ {
		year := i + 2011
		fmt.Printf("Введите количество осадков в %v году ", year)
		fmt.Scan(&rain[i])
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("год: %v осадков: %v\n", 2011+i, rain[i])
	}
}
