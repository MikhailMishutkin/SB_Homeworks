// Функция, возвращающая значение

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomPoint(n int) int { //функция принимает случайное число на вход, и возвращает (второй инт) сгенерированное число в диапозоне от 0 до n
	return rand.Intn(n)
}
func distance(x, y int) { // функция вычисления расстояния между двумя точками
	fmt.Println(x - y)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//x := rand.Intn(10) // рандомная генерация точки в диапазоне от 0 до 10
	//y := rand.Intn(20) // рандомная генерация точки в диапазоне от 0 до 20
	distance(getRandomPoint(10), getRandomPoint(20)) // вызов функции с функциями рандомной генерации в качестве аргументов
	x := getRandomPoint(10)                          //рандомная генерация с помощью функции
	y := getRandomPoint(20)

	distance(x, y)
}
