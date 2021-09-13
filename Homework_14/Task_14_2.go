/* 2. Функция возвращающая несколько значений Что нужно сделать
Напишите программу, которая с помощью функции генерирует 3 случайные точки в двумерном пространстве (две координаты),
 а затем с помощью другой функции преобразует эти координаты по формулам: x = 2 * x + 10, y = -3 * y - 5. */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomPoints(n1, n2 int) (int, int) { //функция принимает случайное два числа на вход, и возвращает два  инта сгенерированное число в диапозоне от 0 до n
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n1), rand.Intn(n2)
}

func getPointsChanged(a int, b int) (x, y int) { // функция принимает на вход два инта, даёт на выход два инта после преобразования
	x = a // присваеваем выходящему значению значение аргумента
	y = b
	x = 2*x + 10
	y = -3*y - 5
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())

	x1, y1 := generateRandomPoints(10, 20)
	x2, y2 := generateRandomPoints(100, 200)
	x3, y3 := generateRandomPoints(1000, 2000)
	fmt.Println(getPointsChanged(x1, y1))
	fmt.Println(getPointsChanged(x2, y2))
	fmt.Println(getPointsChanged(x3, y3))

}
