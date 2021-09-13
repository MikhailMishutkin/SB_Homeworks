//Организация массивов. Жители городов

package main

import "fmt"

func main() {
	var names [5]string
	var citizens [5]int
	for i, _ := range names {
		fmt.Println("Введите название города")
		fmt.Scan(&names[i])
		fmt.Println("Введите количество жителей")
		fmt.Scan(&citizens[i])
	}
	for i, citizen := range citizens {
		if citizen > 100000 {
			fmt.Println(names[i])
		}
	}
}
