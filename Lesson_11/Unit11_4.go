package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "Hello! World"
	b, err := strconv.ParseInt(s[:1], 10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := int(b)
	s = s[:1] + s[6:]
	fmt.Println(s, c)
}
