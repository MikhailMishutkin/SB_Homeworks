/* Задание 1. Сортировка вставками
Что нужно сделать
Напишите функцию, сортирующую массив длины 10 вставками. */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99) * 2
	}
	return slice
}
func sortByInsertion(a []int) []int {
	for i := 1; i < size; i++ { // проходим по массиву с 1
		j := i // второй счётчик начинается всегда с позиции, на до которой массив уже отсортирован
		for j > 0 {
			if a[j-1] > a[j] { // сравниваем значение предыдущего индекса с последующим
				a[j-1], a[j] = a[j], a[j-1] // если он больше, то меняем местами используя синтаксис go
			}
			j = j - 1 // сдвигаем значение индекса влево
		}
	}

	return a
}
func main() {
	fmt.Println("---UNSORT---")
	d := generateSlice(size)
	fmt.Println(d)
	fmt.Println("---SORTED---")
	e := sortByInsertion(d)
	fmt.Println(e)
}
