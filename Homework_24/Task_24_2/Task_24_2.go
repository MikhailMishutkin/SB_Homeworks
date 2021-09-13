/* Задание 2. Анонимные функции
Что нужно сделать
Напишите функцию, которая на вход получает сколько угодно интов,
сортирует их пузырьком и переворачивается (либо сразу сортирует
в обратном порядке, как посчитаете нужным). */

package main

import "fmt"

func inverseBubbleSort(input ...int) []int {
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] < input[j+1] { // обратный порядок сортировки, значения меняются местами если текущее меньше последущего
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}

func main() {
	slice := make([]int, 0, 0) // создаём слайс, куда будем складывать введённые пользователем числа
	for {                      // цикл для ввода значений массива
		var x int
		fmt.Println("Input value integer and 0 to complete array") // ноль для выхода из ввода значений
		fmt.Scanln(&x)
		if x == 0 {
			break
		}
		slice = append(slice, x)
	}

	fmt.Println(inverseBubbleSort(slice...))
}
