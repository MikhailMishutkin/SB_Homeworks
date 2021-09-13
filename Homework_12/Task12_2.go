package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("log.txt") // cоздаём файл куда запишем лог
	if err != nil {
		fmt.Println("Не смогли создать файл, ", err)
		return
	}
	defer file.Close()

	var i int              //счётчик
	var messageUser string // сообщение пользователя
	for messageUser != "выход" {

		fmt.Println("Введите сообщение")
		fmt.Scan(&messageUser)
		i++
		t := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println(i, t, messageUser)
		file.WriteString(fmt.Sprint(i, " ", t, " ", messageUser, "\n"))
	}

}
