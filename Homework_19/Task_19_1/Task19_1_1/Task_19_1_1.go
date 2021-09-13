/* Задание 1. Слияние отсортированных массивов
Что нужно сделать
Напишите функцию, которая производит слияние двух отсортированных массивов длиной
четыре и пять в один массив длиной девять.
Советы и рекомендации
Обратите внимание на размеры массивов. */

package main

import (
	"fmt"
)

const (
	sizeFirst  int = 4 //задаём константами размеры массивов
	sizeSecond int = 5
	sizeUnion  int = 9
)

func uniteArray(input1 [sizeFirst]int, input2 [sizeSecond]int) (thirdArray [sizeUnion]int) {
	var j, k int // вводим счётчики индексов массивов
	for i := 0; i < sizeFirst; i++ {
		a := input1[i]                     // вводим переменную для сравнения значений индексов
		for j := i; j < sizeFirst+1; j++ { //цикл для последовательного сравнения значений последующих индексов второго массива со значением первого массива
			if a > input2[j] {
				thirdArray[k] = a
				k++
				break
			}
			thirdArray[k] = input2[j]
			k++ // увеличение индекса результирующего массива в любом случае
		}

	}
	if j < sizeSecond {
		thirdArray[k] = input2[j+sizeFirst] // присоединение последнего индекса большего массива к результирующему
	}
	return
}

func main() {
	firstArray := [sizeFirst]int{8, 7, 4, 1}
	secondArray := [sizeSecond]int{9, 7, 5, 3, 1}

	fmt.Println(uniteArray(firstArray, secondArray))

}
