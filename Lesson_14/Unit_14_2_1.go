// Функция, возвращающая значение

package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b // return обязателен, когда появляется возвращаемое значение (инт после аргументов), а не только входные аргументы
}
func main() {
	sumAB := sum(20, 30)
	fmt.Println(sumAB)

}
