//Инициализация массивов c range

package main

import "fmt"

func main() {
	var rain [10]int
	for i, _ := range rain {
		year := i + 2011
		fmt.Printf("Введите количество осадков в %v году ", year)
		fmt.Scan(&rain[i])
	}
	for i, r := range rain {
		fmt.Printf("год: %v осадков: %v\n", 2011+i, r)
	}

}
