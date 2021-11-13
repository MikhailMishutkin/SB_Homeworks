// Работа с каналами: паттерн "конвейер" - небуферизированные каналы

package main

import (
	"fmt"
)

func main() {

	fc := putBook()       // первый канал
	sc := deliverBook(fc) // второй канал куда передали первый
	tc := burnBook(sc)    // третий канал, куда передаём значение второго

	//читаем из канала, если без "<-", то выведет адрес канала,
	//а не записанное значение
	fmt.Println(<-tc)
}

func putBook() chan string {
	firstChan := make(chan string) // записываем первое сообщение в первый канал
	go func() {
		firstChan <- "складываю книги"
	}()
	return firstChan
}

func deliverBook(firstChan chan string) chan string {
	secondChan := make(chan string)
	fmt.Println(<-firstChan) // читаем сообщение из первого канала
	go func() {
		secondChan <- "доставляю книги" // записываем второе сообщение во второй канал
	}()
	return secondChan
}

func burnBook(secondChan chan string) chan string {
	thirdChan := make(chan string)
	fmt.Println(<-secondChan)
	go func() {
		thirdChan <- "сжигаю книги"
	}()

	return thirdChan

}
