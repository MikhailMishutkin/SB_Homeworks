package main

import (
	"fmt"
	"strings"
)

/* func main() {
	s := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
	amount := 0
	for _, c := range s {
		d := int(c)
		if d >= 65 && d <= 90 {
			amount++
		}
	}
	fmt.Println(amount)
} */
func main() {
	s := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
	amount := 0
	fmt.Println(len(s))
	fmt.Printf("%T \n", s)
	words := strings.Fields(s)
	fmt.Println(len(words))
	fmt.Printf("%T \n", words)

	for i := 0; i < len(words); i++ {
		word := words[i]
		if word == strings.Title(word) {
			amount++
		}
	}
	fmt.Println("Количество слов с первой заглавной буквой = ", amount)
}

// word := s[:spaceIndex]
// character to ASCII
/* 	firstLet := 'Z' // rune, not string
   	ascii := int(firstLet)
   	fmt.Println(string(char), " : ", ascii) */
/* for len(s) > 0 {
s = strings.Trim(s, " ")
spaceIndex := strings.Index(s, " ")
fmt.Println(spaceIndex)
s = s[:5] */
/* fmt.Println(s)
	a := len(s)
	fmt.Println(a)
	s = strings.Title(s)
	a = len(s)
	fmt.Println(a)
} */
