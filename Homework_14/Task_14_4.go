/* Задание 4. Область видимости переменных
Что нужно сделать
Напишите программу, в которой будет 3 функции, попарно использующие разные глобальные переменные.
Функции, должны прибавлять к поданном на вход числу глобальную переменную и возвращать результат.
Затем вызвать по очереди три функции, передавая результат из одной в другую. */

package main

import "fmt"

const GLOBVAR = 1 // объявляем глобальные переменные типа инт
var a = 5         // подаваемое на вход число
var result int    // возвращаемый результат

func getVar1(a int) (result int) { // создаём функцию складывающую глобальную переменную и число, полученное на вход
	result = GLOBVAR + a
	return
}

func getFuncResult1(f func(int) int) int { // принимаем результат первой функции и присваем его глобальной переменной
	result = f(a)
	return f(a)
}

func getFuncResult2(result int) int { // выводим результат передачи глобальной переменной
	return result
}

func getFuncResult3(result int) int { // проверяем что ничего не меняется
	return result
}

func main() {
	fmt.Println(getVar1(a)) // вызываем по очереди все функции, везде должно быть 6
	fmt.Println(getFuncResult1(getVar1))
	fmt.Println(getFuncResult2(result))
	fmt.Println(getFuncResult3(result))
}
