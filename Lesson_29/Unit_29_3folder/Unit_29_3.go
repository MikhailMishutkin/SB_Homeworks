// Работа с каналами

package main

import (
	"fmt"
)

func main() {

	intChan := make(chan int)

	go func() {
		for i := 0; i < 4; i++ {
			intChan <- i //пишем в канал значения из цикла
		}
		//КАНАЛ ВСЕГДА ЗАКРЫВАЕТ ПИШУЩАЯ ГОРУТИНА,
		//ЕСЛИ ПОПЫТАТЬСЯ ЗАКРЫТЬ КАНАЛ В ЧИТАЮЩЕЙ ТО ПАНИКА!!!
		close(intChan) // закрываем канал после записи значений
	}()

	//первый вариант чтения из канала с использованием логической переменной
	for {
		val, ok := <-intChan //вычитываем значение из канала каждую итерацию цикла
		if !ok {             //если из канала нельзя прочитать, то выходим из цикла
			break
		}
		fmt.Println(val)
	}

	// второй вариант чтения из канала без использования логической переменной
	for val := range intChan {
		fmt.Println(val)
	}
}

/*
func putBook(rchan chan string) {

	fmt.Println("складываю книги")
}

func deliverBook(rchan chan string) {
	fmt.Println("доставляю книги")
}

func burnBook(rchan chan string) {
	fmt.Println("сжигаю книги")

}
*/
