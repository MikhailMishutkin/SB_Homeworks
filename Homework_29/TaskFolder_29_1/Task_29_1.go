// // Задание 1. Конвейер
// Цели задания
// Научиться работать с каналами и горутинами.
// Понять, как должно происходить общение между потоками.
// Что нужно сделать
// Реализуйте паттерн-конвейер:
// Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
// Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
// Произведение: следующая горутина умножает квадрат числа на 2.
// При вводе «стоп» выполнение программы останавливается.
// Советы и рекомендации
// Воспользуйтесь небуферизированными каналами и waitgroup.
// Что оценивается
// Ввод : 3
// Квадрат : 9
// Произведение : 18

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	input := bufio.NewScanner(os.Stdin)

	// первый запрос на ввод
	fmt.Println("Введите число")

	for input.Scan() {
		//проверяем условия ввода на отличие от букв и  0 (возведение в степень не имеет смысла), стоп останавливает программу
		if input.Text() == "стоп" {
			break
		} else {
			fmt.Println("Ввод", input.Text())
			digit, err := strconv.Atoi(input.Text())
			if err != nil {
				fmt.Println(err)
				fmt.Println("Введите число")
				continue
			}
			if digit == 0 {
				fmt.Println("Операция не имеет смысла. Введите число отличное от нуля")
				continue
			}
			//переменная для хранения  квадрата числа в горутине
			mS := makeSquare(digit, &wg)
			fmt.Println("Квадрат: ", mS)
			//  переменная для хранения квадрата числа умноженного на два
			d := double(mS, &wg)
			fmt.Println("Произведение:", d)
		}

		// последующие запросы на ввод
		fmt.Println("Введите число")
	}

	//вызываем функции с аргументом 0 для закрытия горутин и каналов
	makeSquare(0, &wg)
	double(0, &wg)

	fmt.Println("Выполнение программы завершено")
}

func makeSquare(input int, wg *sync.WaitGroup) int {
	// открываю каналы
	fC := make(chan int, 1)
	fC1 := make(chan int, 1)
	fC <- input
	// ставлю вэйтгруппу
	wg.Add(1)
	go func() {
		// завершаю отложенно горутину
		defer wg.Done()
		// закрываю канал
		defer func() {
			close(fC)
			close(fC1)
			fmt.Println("каналы возведения в квадрат закрыты")
		}()

		for value := range fC {
			if value == 0 {
				break
			} else {
				result := value * value
				fC1 <- result
			}
		}
	}()

	return <-fC1
}

func double(fC int, wg *sync.WaitGroup) int {
	sC := make(chan int, 1)
	sC1 := make(chan int, 1)
	sC <- fC
	wg.Add(1)
	go func() {
		// завершаю отложенно горутину
		defer wg.Done()
		// закрываю канал
		defer func() {
			close(sC)
			close(sC1)
			fmt.Println("каналы умножения закрыты")
		}()

		for value := range sC {
			if value == 0 {
				break
			} else {
				result := value * 2
				sC1 <- result
			}
		}
	}()

	return <-sC1
}
