package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hey. we like go. go is cool"
	result := ""
	for len(s) > 0 {
		s = strings.Trim(s, " ")
		spaceIndex := strings.Index(s, " ")
		dotIndex := strings.Index(s, ".")
		minIndex := spaceIndex
		if minIndex == -1 || (dotIndex < spaceIndex && dotIndex != -1) {
			minIndex = dotIndex
		}
		firstWord := s[:minIndex]
		firstWordTitle := strings.Title(firstWord)
		s = strings.Replace(s, firstWord, firstWordTitle, 1)
		if dotIndex >= 0 {
			result += " " + s[:dotIndex+1]
			s = s[dotIndex+1:]
		} else {
			result += " " + s
			s = ""
		}
		fmt.Println(strings.Trim(result, " "))
	}

}
