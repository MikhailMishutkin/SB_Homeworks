/* Урок №6 пакет ioutil
переписать задачи из урока 2 и 3 на пакет ioutil */

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("log.txt") //файл из прошлого урока
	if err != nil {
		panic(err)
	}
	defer f.Close()

	stat, err := f.Stat() //возвращает статистику всего файла
	if err != nil {
		panic(err)
	}

	buf := make([]byte, stat.Size()) // читаем в буфер содержимое файла в байтах
	if _, err := io.ReadFull(f, buf); err != nil {
		panic(err)
	}
	fmt.Printf("%s", buf)
}
