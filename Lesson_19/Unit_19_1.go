// Простые алгоритмы с массивами

package main

import "fmt"

const size = 5

func maximum(input [size]int) int {
	max := input[0] // присваиваем максимум первому элементу
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}

func minimum(input [size]int) int {
	min := input[0] // присваиваем минимум первому элементу
	for i := 0; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}
	return min
}
func main() {
	number := [size]int{10, 49, 48, 20, 39}

	fmt.Println(number[0], number[2]) // обращение к массиву по индексу

	for i := 0; i < len(number); i++ { // len возвращает количество элементов массива
		fmt.Println(number[i])
	}

	fmt.Println(maximum(number), minimum(number))

}
