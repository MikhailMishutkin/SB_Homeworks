/* Задание 1. Слияние отсортированных массивов
Что нужно сделать
Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
Советы и рекомендации
Обратите внимание на размеры массивов. */

package main

import (
	"fmt"
	"sort"
)

const (
	sizeFirst  int = 4 //задаём константами размеры массивов
	sizeSecond int = 5
	sizeUnion  int = 9
)

func uniteArray(input1 [sizeFirst]int, input2 [sizeSecond]int) (thirdArray [sizeUnion]int) { // функция объединенения двух массивов на вход в третий на выход
	for i := range input1 {
		a := input1[i]    // присваиваем переменной значение индекса массива в итерацию
		thirdArray[i] = a // записываем значение третий массив
	}

	for i := range input2 {
		a := input2[i]
		thirdArray[i+len(input1)] = a //записываем значение в третий массив увеличив индекс на размер первого массива
	}
	return
}

func main() {
	firstArray := [sizeFirst]int{1, 2, 3, 4}
	secondArray := [sizeSecond]int{5, 6, 7, 8, 9}

	fmt.Println(uniteArray(firstArray, secondArray))
	s := uniteArray(firstArray, secondArray)
	sort.Ints(s)
	fmt.Println(s)
}
