package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("t.txt")
	if err := os.Chmod("t.txt", 0666); err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer.WriteString("Say hfsd")
	writer.WriteString("\n")
	writer.WriteRune('a')
	writer.WriteString("\n")
	writer.WriteByte(67) // C
	writer.WriteString("\n")
	writer.Write([]byte{65, 66, 67}) //A, B, C
	writer.WriteString("\n")
	//writer.Flush()
}
