//Понятие функций в Go

package main

import "fmt"

func debug() {
	fmt.Println("working")
}

func sum() {
	fmt.Println(5 + 10)
}

func main() {
	debug()
	debug()
	debug()
	debug()
	sum()
}
