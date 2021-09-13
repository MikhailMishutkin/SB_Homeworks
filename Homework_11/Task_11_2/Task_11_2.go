package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "a10 10 20b 20 30c30 30 dd"

	words := strings.Fields(s) //используем функцию филдс для выделения значений из строки от пробелов

	for b := 0; b < len(words); b++ {
		a := words[b]
		i, err := strconv.Atoi(a) // пытаемся конвертировать строку в инт, в случае ошибки - продолжаем
		if err != nil {
			continue
		}
		fmt.Print(i, " ")
	}
}
