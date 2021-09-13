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

func mergingSortedArrays(array1 [sizeFirst]int, array2 [sizeSecond]int) (result [sizeUnion]int) {
	var i1, i2, ir int

	for i1 < sizeFirst && i2 < sizeSecond {
		if array1[i1] < array2[i2] {
			result[ir] = array1[i1]
			i1++
		} else {
			result[ir] = array2[i2]
			i2++
		}
		ir++
	}

	for i1 < sizeFirst {
		result[ir] = array1[i1]
		i1++
		ir++
	}

	for i2 < sizeSecond {
		result[ir] = array2[i2]
		i2++
		ir++
	}
	return
}

func main() {
	fmt.Println("----Слияние сортированных массивов для любой длинны массивов----")
	firstArray := [sizeFirst]int{8, 7, 4, 1}
	secondArray := [sizeSecond]int{9, 7, 5, 3, 1}

	fmt.Println(mergingSortedArrays(firstArray, secondArray))

}
