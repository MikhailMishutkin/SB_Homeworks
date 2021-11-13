// Знакомство с WaitGroup и горутинами

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup //структура для синхронизации горутин
	//с помощью метода Add указываем сколько горутин
	//будем использовать/хотим создать (функций 3, но main сама горутина)
	wg.Add(2)

	go putBook(&wg) // передача по указателю обязательно!!!
	go deliverBook(&wg)

	// ставим метод Wait перед последней функцией, которую хотим испольнить - он ждёт испольнения всех остальных
	// wg, т.е. счётчик wg становиться равным нулю
	wg.Wait()
	burnBook()
}

func putBook(wg *sync.WaitGroup) {
	defer wg.Done() // ставим окончание функции и уменьшения счётчика wg с помощью defer или в конце
	fmt.Println("складываю книги")
}

func deliverBook(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("доставляю книги")
}

func burnBook() {
	fmt.Println("сжигаю книги")

}
