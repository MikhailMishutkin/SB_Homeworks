package main

import (
	"fmt"
	"strings"
)

func main() {

	var s string
	s = "Golang is programming language"
	a := len(s)
	s1 := strings.ReplaceAll(s, "a", "")
	b := len(s1)
	fmt.Print("количество 'a' в предложении равно ", a-b)
}
