package main

import "fmt"

func first(a int) {
	fmt.Println(a)
}
func second(s int) {
	fmt.Println(s)
}
func big(first func(int), second func(int)) {
	first(1)
	second(2)
}

func main() {
	big(second, first)
}
