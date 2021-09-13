//Определенное и неопределенное поведение в функции

package main

import "fmt"

func division(a, b int) {
	if b == 0 { //т.к. на ноль делить нельзя, условие необходимо прописать в теле функции
		fmt.Println("UNKNOWN RESULT")
	} else {
		fmt.Println(a / b)
	}
}

func notDebug(s string) {
	if s == "debug" {
		fmt.Println("NOT DEBUG")
	} else {
		fmt.Println(s)
	}
}

func checkPassword(password string) {
	if len(password) < 6 {
		fmt.Println("pass is to short")
	} else {
		fmt.Println("success")
	}
}

func printPointer(a *int) {
	if a != nil { // избегаем ситуации с неопределенным указателем в мэйн
		fmt.Println(*a)
	}

}

func main() {
	division(3, 3)
	division(6, 3)
	division(6, 0)

	notDebug("some asdf")
	notDebug("debug")

	checkPassword("password")
	checkPassword("pass")

	var a *int // неопределенный указатель
	b := 10
	c := &b
	printPointer(c)
	printPointer(a) // результат не будет печататься, т.к. а не проходит по условию внутри функции
}
