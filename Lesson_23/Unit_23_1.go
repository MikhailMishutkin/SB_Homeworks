// разбор домашней работы 22

package main

import (
	"fmt"
)

const size = 10

func find(a [size]int, n int) (index int) { // поиск для второй задачи
	index := -1

	min := 0
	max := size - 1
	for max >= min {
		middle := (max + min) / 2
		if a[middle] == n {
			index = middle
			break
		} else if a[middle] > n {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	for index > 0 && a[index-1] == n {
		index--
	}
	return
}
func main() {
	//rand.Seed(time.Now().UnixNano()) // для первой задачи

	var numbers [size]int
	for i := 0; i < size; i++ {
		//	 numbers[i] = rand.Intn(10*size) // Для первой задачи
		fmt.Scan(&numbers[i])
	}

	fmt.Printf("%+v\n", numbers)

	n := 0
	fmt.Scan(&n)

	count := 0 // первая задача
	found := false

	for i := 0; i < size; i++ {
		if found {
			count++
		} else if numbers[i] == n {
			found = true
		}
	}
	fmt.Println("Цифры после вашей: ", count)

	index := find(numbers, n) // вывод резултата решения второй задачи
	fmt.Println("Индекс ", index)
}
