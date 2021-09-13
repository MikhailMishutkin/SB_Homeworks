package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return
	}
	defer f.Close()
	fileSize, err := f.Stat() // вводим переменную, которая будет содержать информацию о файле
	if err != nil {
		panic(err)
	}
	// fmt.Println(fileSize.Size()) // проверяем размер файла с помощью Size
	buf := make([]byte, fileSize.Size())
	if _, err := io.ReadFull(f, buf); err != nil {
		fmt.Println("Не смогли прочитать последовательность байтов из файла", err)
		return
	}

	fmt.Println(string(buf))
}

/* buf := make([]byte, 512)
_, err = f.Read(buf)
if err != nil {
	fmt.Println("Не смогли прочитать последовательность байтов из файла", err)
	return
} */
