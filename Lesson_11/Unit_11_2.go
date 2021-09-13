package main

import (
	"fmt"
	"strconv"
)

func main() {
x:
	fmt.Println("Введите число")
	var a string
	fmt.Scan(&a)

	fmt.Println("Укажите систему исчисления для перевода: 2 или 16")
	var b string
	fmt.Scan(&b)

	i, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
		goto x
	}

	m := int64(i)

	if b == "2" {
		a := strconv.FormatInt(m, 2)
		fmt.Println(a)
	} else if b == "16" {
		a := strconv.FormatInt(m, 16)
		fmt.Println(a)
	} else {
		fmt.Println("Вы ввели неверное значение")
		goto x
	}
}
