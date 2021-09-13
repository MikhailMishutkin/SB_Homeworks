//Проверка на сортированный массив. Метод Монте-Карло

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
		a[i] = /* 10*i + */ rand.Intn(10)
	}
	fmt.Println(checkSorted(a))
	fmt.Println(checkSorted1(a))
}

func checkSorted(a [n]int) (result bool) {
	result = true
	for i := 0; i < n-1; i++ {
		if a[i+1] < a[i] {
			result = false
			fmt.Println("Определили методом перебора")
			break
		}
	}
	return
}

//Метод Монте-Карло

func checkSorted1(a [n]int) (result bool) {
	result = true
	firstIndex := rand.Intn(n / 3)
	secondIndex := rand.Intn(n/3) + n/3
	thirdIndex := rand.Intn(n/3) + 2*n/3
	if a[firstIndex] > a[secondIndex] || a[secondIndex] > a[thirdIndex] || a[firstIndex] > a[thirdIndex] {
		result = false
		fmt.Println("Определили методом Монте-Карло")
		return
	}
	return
}
