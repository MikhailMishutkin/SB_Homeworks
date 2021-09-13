package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("log.txt") // открываем файл прошлого задания
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return
	}
	defer f.Close()

	resultBytes, err := ioutil.ReadAll(f) // читаем файл пакетом ioutil
	if err != nil {
		panic(err)
	}
	fmt.Println("Сохранённый лог: \n")
	fmt.Println(string(resultBytes)) // выводим результат чтения файла строкой
}
