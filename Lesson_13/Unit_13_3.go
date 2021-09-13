//Аргументы функции и способы их передачи

package main

import "fmt"

func sum() { // функции без входящих аргументов с заданным результатом на выход
	fmt.Println(5 + 10)
}

func sum1() {
	fmt.Println(2 + 3)
}

func sumTwoValues(a, b int) { // функции с входящими аргументами
	fmt.Println(a + b)
}

func sumTwoDiffValues(a uint64, b int32) {
	fmt.Println(a + uint64(b))
}

func debug(s string) {
	fmt.Println(s)
}
func main() {

	sum()
	debug("first")
	sum1()
	debug("second")
	sumTwoValues(3, 4)
	debug("sum")
	sumTwoDiffValues(uint64(100), int32(101))
	debug("sumTwoDiffValues")
}
