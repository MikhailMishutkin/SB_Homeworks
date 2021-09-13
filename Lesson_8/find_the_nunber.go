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
	for {
		var answer int
		for {
			_, _ = fmt.Scan(&answer)
			if answer < 1 || answer > 100 {
				fmt.Println("Число должно быть в диапозоне от 1 до 100")
			} else {
				break
			}
		}
		if answer == n {
			fmt.Println("Вы угадали!!!")
			break
		} else if answer < n {
			fmt.Println("Загаданное число больше")
		} else {
			fmt.Println("Загаданное число мешьше")
		}
	}

}
