package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	/* file, err := os.Create("log.txt") // cоздаём файл куда запишем лог
	if err != nil {
		fmt.Println("Не смогли создать файл, ", err)
		return
	}
	defer file.Close() */

	var b bytes.Buffer // вводим переменную для буфера

	var i int              //счётчик cтрок
	var messageUser string // сообщение пользователя
	for messageUser != "выход" {

		fmt.Println("Введите сообщение")
		fmt.Scan(&messageUser)
		i++
		t := time.Now().Format("2006-01-02 15:04:05") // вводим переменную, где записана дата и время сообщения
		fmt.Println(i, t, messageUser)
		line := fmt.Sprintf("%d %v %s\n", i, t, messageUser) // записываем в буфер строку со всеми данными
		if _, err := b.WriteString(line); err != nil {
			panic(err)
		}
	}
	fiName := "log.txt"
	if err := ioutil.WriteFile(fiName, b.Bytes(), 0666); err != nil { // создаем файл с помощью ioutil
		panic(err)
	}
	fi, err := os.Open(fiName)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	resultBytes, err := ioutil.ReadAll(fi) // читаем весь файл
	if err != nil {
		panic(err)
	}
	fmt.Println("Сохранённый лог: \n")
	fmt.Println(string(resultBytes))
}
