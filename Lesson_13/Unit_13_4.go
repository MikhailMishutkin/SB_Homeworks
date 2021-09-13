//передача по ссылке, использование указателей

package main

import "fmt"

func makeItZero(x int) {
	x = 0
}

func makeItZeroPointer(x *int) {
	*x = 0
}

func pIncrement(x *int) {
	*x++
}

func main() {
	a := 10
	makeItZero(a) // возвращает значение 10, а не 0, т.к. обращается напрямую к "a"
	fmt.Println(a)
	makeItZeroPointer(&a) // возвращает результат выполения функции, т.е. 0, т.к. содержит указатель
	fmt.Println(a)

	b := 10
	c := &b       //указываем, что переменная работает с указателем
	pIncrement(c) // нет амперсанта, указали заранее
	fmt.Println(b)

}
