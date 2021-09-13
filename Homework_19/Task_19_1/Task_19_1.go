/* Задание 1. Слияние отсортированных массивов
Что нужно сделать
Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
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
	var i, j, k int
	for i < sizeFirst && j < sizeSecond {
		if input1[i] >= input2[j] {
			thirdArray[k] = input1[i]
			k++
			i++
		} else {
			thirdArray[k] = input2[j]
			k++
			j++
		}

	}
	for i < sizeFirst {
		thirdArray[k] = input1[i]
		k++
		i++
	}
	for j < sizeSecond {
		thirdArray[k] = input2[j]
		k++
		j++
	}
	return
}

func main() {
	fmt.Println("----Слияние сортированных массивов----")
	firstArray := [sizeFirst]int{8, 7, 4, 1}
	secondArray := [sizeSecond]int{9, 7, 5, 3, 1}

	fmt.Println(uniteArray(firstArray, secondArray))

}
