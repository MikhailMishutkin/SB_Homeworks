// Строки в контексте слайсов

package main

import (
	"fmt"
	"reflect"
	"strings"
)

func debugSlice1(input []byte) {
	fmt.Printf("data: %+v\n", input)
	fmt.Printf("len: %+v\n", len(input))
	fmt.Printf("cap: %+v\n", cap(input))
}

func main() {
	exampleString := "empty value"
	for i, v := range exampleString {
		fmt.Println(i, string(v))
	}

	// string <=> byte
	debugSlice1([]byte(exampleString))

	//strings
	a := "airst"
	b := "first"
	fmt.Println(strings.Compare(a, b)) // сравниватся значения байт для букв
	//compare a==b, a>b
	fmt.Println([]byte("a"), []byte("f"), []byte("z"))

	//index() - находит значение 2 в значении 1, если нет, то результат -1
	index := strings.Index(a, b)
	if index != -1 { // оборачиваем в условие -1
		//todo
	}

	// split() - убирает второе значение из первого
	e := "127.0.0.1"
	f := strings.Split(e, ".") // убирает точку из строки e
	fmt.Println(f)
	fmt.Printf("%+v\n", reflect.TypeOf(f))
	for _, v := range f {
		fmt.Println(v)
	}

}
