package main

import (
	"fmt"
)

func main() {
	s := "(a)((b())())())("
	count := 0
	for len(s) > 0 {
		if s[:1] == "(" {
			count++
		} else if s[:1] == ")" {
			count--
		}
		s = s[1:]
		if count < 0 {
			break
		}
	}
	if count == 0 {
		fmt.Println("Все хорошо")
	} else {
		fmt.Println("Не совпадают")
	}
}
