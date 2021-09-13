//Поиск в сортированном массиве

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10 // количество элементов массива

func main() {
	rand.Seed(time.Now().UnixNano()) // генерация массива
	var a [n]int
	for i := 0; i < n; i++ {
		a[i] = 10*i + rand.Intn(10)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%v\t", a[i])
	}
	fmt.Println()
	// Ищем число находящееся в массиве
	value := a[rand.Intn(n)]
	index := find(a, value)
	fmt.Printf("Индекс: %v\n", index)
	//Ищем число НЕ находящееся в массиве
	value = 20 * n
	fmt.Println(value)
	index = find(a, value)
	fmt.Printf("Индекс: %v\n", index)
}

func find(a [n]int, value int) (index int) {
	index = -1
	min := 0
	max := n - 1
	for max >= min {
		middle := (max + min) / 2
		if a[middle] == value {
			index = middle
			break
		} else if a[middle] > value {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	return
}
