package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 8)
	if _, err = io.ReadFull(f, buf); err != nil {
		fmt.Println("Не смогли прочитать последовательность байтов из файла", err)
		return
	}
	/* buf := make([]byte, 512)
	_, err = f.Read(buf)
	if err != nil {
		fmt.Println("Не смогли прочитать последовательность байтов из файла", err)
		return
	} */
	fmt.Println(string(buf))

}
