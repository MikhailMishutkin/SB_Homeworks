//Panic(ошибка runtime) при работе с массивами

package main

import "fmt"

const size = 5 // длина массива только константа, иначе ошибка на уровне компилятора

func newNumbers(input [size]int) {
	for i := 0; i < 6; i++ { // ошибка в runtime(во время исполнения) i > size, другие функции выполнятся
		if i < size { // для избежания ошибки в  runtime можно обернуть функцию в условие
			fmt.Println(input[i])
		} else {
			continue
		}
	}
}

func main() {
	number := [size]int{10, 49, 48, 20, 39}

	/* fmt.Println(number[0], number[2]) // обращение к массиву по индексу

	for i := 0; i < len(number); i++ { // len возвращает количество элементов массива
		fmt.Println(number[i])
	}
	*/

	newNumbers(number)
}
