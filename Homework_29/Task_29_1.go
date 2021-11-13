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
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите число")
	for input.Scan() {
		if input.Text() == "стоп" {
			break
		} else {
			fmt.Println("Ввод", input.Text())
			mS := makeSquare(input.Text()) //канал для  квадрата числа в горутине
			fmt.Println("Квадрат: ", mS)
			d := double(mS) // канал куда передаём умноженное на два
			fmt.Println("Произведение:", d)
		}

		fmt.Println("Введите число")
	}

}

func makeSquare(input string) int {
	fC := make(chan int)

	go func() {
		a, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		a = a * a
		fC <- a
	}()
	return <-fC
}

func double(fC int) int {
	sC := make(chan int)
	go func() {
		fC = fC * 2
		sC <- fC
	}()
	return <-sC
}
