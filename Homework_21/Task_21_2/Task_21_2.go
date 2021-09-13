/* Задание 2. Анонимные функции
Что нужно сделать
Напишите функцию, которая на вход принимает функцию вида A func (int, int) int,
а внутри оборачивает и вызывает её при выходе (через defer).
Вызовите эту функцию с тремя разными анонимными функциями A.
Тела функций могут быть любыми, но главное, чтобы все три выполняли разное действие. */

package main

import "fmt"

func justCallFuncA(A func(a int, b int) int, a int, b int) (result int) {
	defer func() {
		result = A(a, b)
	}()
	result = 0
	return
}

func main() {

	fmt.Println(justCallFuncA(func(a int, b int) int { return a - b }, 40, 5))
	fmt.Println(justCallFuncA(func(a int, b int) int { return a / b }, 40, 5))
	fmt.Println(justCallFuncA(func(a int, b int) int { return a * b }, 40, 5))
}

/* package main

import "fmt"

func justCallFuncA(A func(a int, b int) int) {
	a := 40
	b := 5
	defer fmt.Println(A(a, b))
}

func main() {

	justCallFuncA(func(a int, b int) int { return a - b })
	justCallFuncA(func(a int, b int) int { return a / b })
	justCallFuncA(func(a int, b int) int { return a * b })
} */
