// Функция, возвращающая нескольо значений

package main

import (
	"fmt"
	"math/rand"
)

func getRandomnumberAndString(n int) (int, string) { // на вход только число, на выход число и строка
	return rand.Intn(n), "HELLO" // число генерируется, строка задана сразу
}
func main() {
	fmt.Println(getRandomnumberAndString(10)) // выводим значение функции при n от 0 до 10
	x, str := getRandomnumberAndString(10)    // присваиваем значения функции переменным
	fmt.Println("INT", x)                     // выводим построчно оба значения
	fmt.Println("STR", str)
}
