package main

import "fmt"

const (
	firstSpartian uint16 = 1 + iota
	secondSpartian
	thirdSpartian
	fourthSpartian
	fifthSpartian
)

func main() {
	fmt.Printf("спартанец номер %d \n", firstSpartian)
	fmt.Printf("спартанец номер: %d \n", secondSpartian)
	fmt.Printf("спартанец номер: %d \n", thirdSpartian)
	fmt.Printf("спартанец номер: %d \n", fourthSpartian)
	fmt.Printf("спартанец номер: %d \n", fifthSpartian)
}
