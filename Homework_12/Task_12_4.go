package main

import (
	"fmt"
	"os"
)

func main() {
	text := "Hello!"
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Println(file.Name())
	err = file.Chmod(0444)
	if err != nil {
		fmt.Println(err)
	}

	var message string
	fmt.Println("Введите текст")
	fmt.Scan(&message)

	file.WriteString(message)
	file.WriteString("\n")
	file.WriteString(text)

}
