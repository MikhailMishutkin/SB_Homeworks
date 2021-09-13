// работа с массивом рун

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	str := "привет"
	for i := 0; i < len(str); i++ { // выводит юникод
		fmt.Println(str[i])
	}
	fmt.Println()

	strRune := []rune(str)
	for i := 0; i < len(strRune); i++ { // выводит байты
		fmt.Println(strRune[i])
	}
	fmt.Println()

	for _, r := range strRune {
		fmt.Println(r, string(r)) // вывод букв построчно с байтами

	}
	fmt.Println()

	b := []byte("hello, привет")
	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b) // возвращает в обратном порядке руну и количество используемых байт
		fmt.Printf("%c, %v\n", r, size)
		b = b[:len(b)-size]
	}
	fmt.Println()

	b = []byte("hello, привет")
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b) // возвращает в обычном порядке руну и количество используемых байт
		fmt.Printf("%c, %v\n", r, size)
		b = b[size:]
	}

	message := "sometext"
	for i := 0; i < len(message); i++ { // код Цезаря
		c := message[i]
		if c >= 'a' && c <= 'z' { //сравниваем букву в итерации с руной
			c = c + 13   // меняем положение сдвигом
			if c > 'z' { //если больше алфавита, вычитаем 26
				c = c - 26
			}
			fmt.Printf("%+c", c)
		}

	}
}
