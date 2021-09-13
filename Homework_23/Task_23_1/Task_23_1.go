/* Задание 1. Чётные и нечётные
Что нужно сделать
Напишите функцию, которая принимает массив чисел, а возвращает два массива:
один из чётных чисел, второй из нечётных.  */

package main

import "fmt"

const size = 10

func divideArray(a [size]int) (evenArr [size / 2]int, notEvenArr [size / 2]int) { //на вход один массив, на выход два массива пополам
	k, j := 0, 0 // счётчики выходных массивов
	for i, r := range a {
		if r%2 == 0 {
			evenArr[k] = a[i]
			k++ // при попадании в условие увеличивает счётчик на 1
		} else {
			notEvenArr[j] = a[i]
			j++
		}
	}
	return
}

func main() {
	arrayToGet := [size]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	even, notEven := divideArray(arrayToGet)

	fmt.Println(even)
	fmt.Println(notEven)
}
