//defer и анонимные функции

package main

import "fmt"
 func iAcceptFunctions(x, y, z int, F func(int, int)int) int { на вход три числа и функция
	return x + F(y, z)
 }

 func iAmUsingDefer() {
	// defer fmt. Println("exiting function") // вызывается после тела функции
	defer fun() { // в defer можно подавать готовые функции (анонимные)
		fmt.Println(1)
		fmt.Println(2)
	}() // пустые скобки говорят о вызове обёрнутой функции с пустым количеством аргументов
	fmt.Println("before defer")
 }

func main() {
	//f:=func (x, y int) int {return x + y} // пример анонимной функции
	fmt.Println(iAcceptFunctions(10, 20, 30, func(y, z int) int {return y+z})) // входящая функция анонимная, даёт результат в теле вызова первичной
	iAmUsingDefer()
	fmt.Println("DONE") //после defer, т.к. defer только в функции
}