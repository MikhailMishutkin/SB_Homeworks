/* Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться)
Что нужно сделать
Заполните упорядоченный массив из 12 элементов и введите число.
Необходимо реализовать поиск первого вхождения заданного числа в массив.
Сложность алгоритма должна быть минимальная.
Что оценивается
Верность индекса.
При вводе массива 1 2 2 2 3 4 5 6 7 8 9 10 и вводе числа 2 программа должна вывести индекс 1.
*/
package main

import "fmt"

const n = 12

var a [n]int

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

func compareNearIndex(a [n]int, value int) int { // Сравниваем стоящие рядом члены массива и присваеваем значение, если меньший индекс равен действующему
	for i := 0; i < value; i++ {
		if a[value] == a[value-1] {
			value = value - 1
		}
	}
	return value
}

func main() {
	for i, _ := range a { // запрос чисел массива у пользователя
		number := i + 1
		fmt.Printf("Введите %v-ю цифру \n", number)
		fmt.Scan(&a[i])
	}
	fmt.Println("Введите повторяющееся число из массива")
	var repeatNum int
	fmt.Scan(&repeatNum)
	s := find(a, repeatNum)
	fmt.Println(compareNearIndex(a, s))
}
