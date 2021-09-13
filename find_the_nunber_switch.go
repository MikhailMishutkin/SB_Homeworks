package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(101)
	fmt.Println("Введите число от 1 до 100")
outerLoop:
	for {
		var answer int
	innerLoop:
		for {
			_, _ = fmt.Scan(&answer)
			switch {
			case answer < 1 || answer > 100:
				fmt.Println("Число должно быть в диапозоне от 1 до 100")
			default:
				break innerLoop
			}
			switch {
			case answer == n:
				fmt.Println("Вы угадали!!!")
				break outerLoop
			case answer < n:
				fmt.Println("Загаданное число больше")
			default:
				fmt.Println("Загаданное число мешьше")
			}
		}
	}
}
