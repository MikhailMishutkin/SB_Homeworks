package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите логин")
	var log string
	fmt.Scan(&log)
	fmt.Println("Введите пароль")
	var pass string
	fmt.Scan(&pass)
	fmt.Println("Введите возраст")
	var age string
	fmt.Scan(&age)

	file, err := os.Create("credentials.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Ваш логин: %s \n", log))
	b.WriteString(fmt.Sprintf("Ваш пароль: %s \n", pass))
	b.WriteString(fmt.Sprintf("Ваш возраст: %s \n", age))
	_, err = file.Write(b.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println("===================")
	// fmt.Println("log, ",", "вы успешно зашли!" )
}
