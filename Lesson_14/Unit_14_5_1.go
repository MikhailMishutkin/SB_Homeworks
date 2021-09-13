// Работа с общими для всех функций переменными
package main

import "fmt"

const c = 100 // объявление глобальной переменной,

func f1(a int) int { //функция с переменной на вход и значением инт на выход
	return a - c
}
func f2(b int) int { //функция с переменной на вход и значением инт на выход
	return b + c
}

func main() {
	fmt.Println(f1(10))
	fmt.Println(f2(30))
}
