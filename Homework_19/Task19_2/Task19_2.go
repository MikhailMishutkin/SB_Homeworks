/* Задание 2. Сортировка пузырьком
Что нужно сделать
Отсортируйте массив длиной шесть пузырьком. */

package main

import "fmt"

func buble(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}

func main() {
	fmt.Println("===Сортировка пузырьком===")
	getArray := []int{2, 8, 6, 4, 7, 1}
	fmt.Println(getArray)
	fmt.Println(buble(getArray))
}
