// Функция, возвращающая нескольо значений

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomPoint(n int) int { //функция принимает случайное число на вход, и возвращает (второй инт) сгенерированное число в диапозоне от 0 до n
	return rand.Intn(n)
}
func getRandom2DPoint(n1, n2 int) (int, int) { //функция принимает случайное два числа на вход, и возвращает два  инта сгенерированное число в диапозоне от 0 до n
	return rand.Intn(n1), rand.Intn(n2)
}
func distance(x, y int) { // функция вычисления расстояния между двумя точками
	fmt.Println(x - y)
}

func distanceNew(x1, y1, x2, y2 int) { // функция вычисления расстояния между  точками в 2-х мерном пространстве
	fmt.Println(x1 - y1)
	fmt.Println(x2 - y2)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	distance(getRandomPoint(10), getRandomPoint(20)) // вызов функции с функциями рандомной генерации в качестве аргументов
	x1, y1 := getRandom2DPoint(10, 20)               // создание первой точки, значения n - произвольные
	x2, y2 := getRandom2DPoint(100, 200)             // создание второй точки
	distanceNew(x1, y1, x2, y2)
}
