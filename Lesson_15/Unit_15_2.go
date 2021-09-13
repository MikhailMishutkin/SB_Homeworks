// Индексация массивов

package main

import "fmt"

func main() {
	// var a [5]int
	a := [...]int{10, 20, 30, 40, 50}
	a[4] = 75
	//fmt.Println(a)
	fmt.Println(a[4])
}
