/* Задание 1. Расчёт по формуле
Что нужно сделать
Напишите функцию, производящую следующие вычисления.
S = 2 × x + y ^ 2 − 3/z, где x — int16, y — uint8, a z — float32.
Тип S должен быть во float32.
Обратите внимание, к какому типу надо привести конечный результат. */

package main

import (
	"fmt"
)

func main() {
	var x int16
	var y uint8
	var z float32

	/* x = 25000
	y = 50
	z = 3 */

a1:
	fmt.Println("Введите первое число для вычисления в диапазоне от -32768 до 32767")
	w := float32(x)
	fmt.Scan(&w)
	fmt.Println(x)
	if w > 32767 || w < -32768 {
		fmt.Println("Введено число вне диапазона")
		goto a1
	}
a2:
	fmt.Println("Введите второе число для вычисления в диапазоне от 0 до 255")
	r := int(y)
	fmt.Scan(&r)
	if r > 255 || r < 0 {
		fmt.Println("Введено число вне диапазона")
		goto a2
	}

	fmt.Println("Введите третье число для вычисления")
	fmt.Scan(&z)

	s := func(a int16, b uint8, c float32) float32 { return float32(a)*2 + float32(b)*float32(b) - 3/c }

	fmt.Println(s(x, y, z))
}
