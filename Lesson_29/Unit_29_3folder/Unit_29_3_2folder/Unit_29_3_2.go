// Работа с каналами: паттерн "конвейер" - Буферизированные каналы

package main

import (
	"fmt"
)

func main() {

	bufChan := make(chan int, 3) // вместимость 3

	//записываем два значения
	bufChan <- 1
	bufChan <- 1

	fmt.Println(len(bufChan), cap(bufChan)) // выводим длинну и вместимость канала: 2 и 3
}
