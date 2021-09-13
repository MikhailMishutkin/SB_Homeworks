package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello World 1"
	for strings.Contains(s, " ") {
		spaceIndex := strings.Index(s, " ")
		//fmt.Println(spaceIndex)
		fmt.Println(s[:spaceIndex])
		s = s[spaceIndex+1:]
	}
	fmt.Println(s)

}
