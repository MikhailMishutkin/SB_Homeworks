package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "a10 10 20b 20 30c30 30 dd"

	words := strings.Fields(s)

	for b := 0; b < len(words); b++ {
		a := words[b]
		i, err := strconv.Atoi(a)
		if err != nil {
			res64, err := strconv.ParseInt(a, 16, 32)
			if err != nil {
				panic(err)
			}
			fmt.Println(a, "эта часть строки преобразуется в 16-ричную систему, результат: ", res64)
		} else {
			fmt.Println(i, " эта часть строки приводится к целому числу")
		}
	}
}
