/* Дополнительно задание со звездочкой*
Написать программу, которая на вход принимала бы интовое
 число и для него генерировала бы все возможные
комбинации круглых скобок, т.е. на вход приходит число 3
на выходе: ["((()))","(()())","(())()","()(())","()()()"]
на вход 1
на выходе: ["()"] */

package main

import (
	"fmt"
	"strings"
)

func generate(open, closing int, queue string, result *[]string) { // резалт это указатель на массив
	if open == 0 {
		*result = append(*result, queue+strings.Repeat(")", closing))
		return
	}
	if closing > open {
		generate(open, closing-1, queue+")", result)
	}
	generate(open-1, closing, queue+"(", result)
}
func main() {
	result := make([]string, 0)
	n := 2
	generate(n, n, "", &result)
	fmt.Println(result)
}
