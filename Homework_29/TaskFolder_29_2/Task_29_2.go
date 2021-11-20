// Задание 2. Graceful shutdown
// Цель задания
// Научиться правильно останавливать приложения.
// Что нужно сделать
// В работе часто возникает потребность правильно останавливать приложения. Например, когда наш сервер обслуживает соединения,
//а нам хочется, чтобы все текущие соединения были обработаны и лишь потом произошло выключение сервиса. Для этого существует паттерн graceful shutdown.
// Напишите приложение, которое выводит квадраты натуральных чисел на экран, а после получения сигнала ^С обрабатывает этот сигнал, пишет «выхожу из программы» и выходит.
// Советы и рекомендации
// Для реализации данного паттерна воспользуйтесь каналами и оператором select с default-кейсом.
// Что оценивается
// Код выводит квадраты натуральных чисел на экран, после получения ^С происходит обработка сигнала и выход из программы.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

const workersPool = 2

type NaturalNumbers struct {
	Num int
}

func main() {
	// создаём слайс структуры
	var row []NaturalNumbers
	var wg sync.WaitGroup

	// наполняем слайс натуральными числами
	for i := 1; i < 100; i++ {
		data := NaturalNumbers{Num: i}
		row = append(row, data)
	}

	poolWork(row, &wg)

}

//функция принимает на вход слайс натуральных чисел и вэйт группу
func poolWork(row []NaturalNumbers, wg *sync.WaitGroup) {
	// канал для вывода квадратов натуральных чисел
	ch := make(chan NaturalNumbers, workersPool)
	// канал для остановки выполнения программы пользователем
	interruptChan := make(chan os.Signal, 1)

	//цикл запускающий горутины
	for i := 1; i < workersPool; i++ {
		wg.Add(1)
		go func() {
			// цикл для чтения из канала
			for data := range ch {

				result := data.Num * data.Num

				signal.Notify(interruptChan, os.Interrupt)

				// в случае заполнения канала для остановки пользователем срабатывает case, иначе default
				select {
				case <-interruptChan:
					close(ch)
					wg.Done()
					fmt.Println("Программа завершена")
					os.Exit(0)
				// выводим квадраты чисел с небольшой задержкой
				default:
					time.Sleep(1000 * time.Millisecond)
					fmt.Println("Квадрат", result/data.Num, "равен", result)
				}
			}

		}()
	}
	// цикл для наполнения канала из слайса
	for i, _ := range row {
		ch <- row[i]
	}
}
