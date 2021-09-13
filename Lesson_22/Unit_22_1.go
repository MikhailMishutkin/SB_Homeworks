//Поиск в несортированном массиве

package main

import (
	"math/rand"
	"time"
)

const n = 10 // количество элементов массива

func main() {
	rand.Seed(time.Now().UnixNano()) // генерация массива
	var a [n]int
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(10 * n)
	}
	// Ищем число находящееся в массиве
	value : = a[n/2]
	index := find(a, value)
	fmt.Println("Индекс: %v\n", index)
	// ищем число не находящеесяс в массиве
	value := n*20
	index = find(a, value)
	fmt.Println("Индекс: %v\n", index)
}

func find(a [n]int, value int) (index int) { // поиск в массиве с присвоением результата поиска значению 1/-1 (найдено/не надйено)
	index = -1 // не найдено по умолчанию
	for i := 0; i < n; i++ {
		if a[i] == value { // условие прерывания цикла, если число найдено в массиве
			index = i
			break
		}
	}
	return
}
