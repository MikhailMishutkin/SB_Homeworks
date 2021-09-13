package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello  World "
	s = strings.Trim(s, " ")
	for strings.Contains(s, " ") {
		spaceIndex := strings.Index(s, " ")
		fmt.Println(s[:spaceIndex])
		s = s[spaceIndex+1:]
		s = strings.Trim(s, " ")
	}
	fmt.Println(s)

}
