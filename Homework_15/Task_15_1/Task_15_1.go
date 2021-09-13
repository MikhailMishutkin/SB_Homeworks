/* Подсчёт чётных и нечётных чисел в массиве
Что нужно сделать
Разработайте программу, позволяющую ввести 10 целых чисел, а затем вывести из них количество
 чётных и нечётных чисел. Для ввода и подсчёта используйте разные циклы.
Что оценивается
Для введённых чисел 1, 1, 1, 2, 2, 2, 3, 3, 3, 4 программа должна вывести: чётных — 4, нечётных — 6 */

package main

import "fmt"

var a [10]int
var even int
var notEven int

func evenNotEven(a []int) (even, notEven int) {
	for _, r := range a {
		if r%2 == 0 {
			even++
		} else {
			notEven++
		}
	}
	return
}

func main() {
	for i, _ := range a {
		number := i + 1
		fmt.Printf("Введите %v-ю цифру \n", number)
		fmt.Scan(&a[i])
	}
	even, notEven := evenNotEven(a[0:10])
	fmt.Printf("Чётных - %v, нечётных - %v", even, notEven)
}
