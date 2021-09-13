package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Не смогли создать файл, ", err)
		return
	}
	defer file.Close()
	rand.Seed(time.Now().UnixNano()) // активировали рандомайзер
	n := rand.Intn(101)              // задали числовой отрезок для ранда
	fmt.Println("Введите число от 1 до 100")
	file.WriteString("Введите число от 1 до 100 \n")
	for {
		var answer int
		for {
			_, _ = fmt.Scan(&answer)
			file.WriteString(fmt.Sprintf("Введено число %d \n", answer))
			if answer < 1 || answer > 100 {
				fmt.Println("Число должно быть в диапозоне от 1 до 100")
				file.WriteString("Число должно быть в диапозоне от 1 до 100 \n")
			} else {
				break
			}
		}
		if answer == n {
			fmt.Println("Вы угадали!!!")
			file.WriteString("Число угадано")
			break
		} else if answer < n {
			fmt.Println("Загаданное число больше")
			file.WriteString("Загаданное число больше \n")
		} else {
			fmt.Println("Загаданное число мешьше")
			file.WriteString("Загаданное число меньше \n")
		}
	}
	// file.WriteString(log)
	//fmt.Println(file.Name())
}
