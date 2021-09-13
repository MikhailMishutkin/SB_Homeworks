// добавление в срезы и обход

package main

import "fmt"

func debugSlice(input []int) {
	fmt.Printf("data: %+v\n", input)
	fmt.Printf("len: %+v\n", len(input))
	fmt.Printf("cap: %+v\n", cap(input))
}

func main() {
	//testSlice := make([]int, 0, 0)

	//testSlice = append(testSlice, 1) // первый аргумент это КУДА добавить, второй - ЧТО добавить, возвращает значение
	/* fmt.Printf("%+v\n", testSlice)
	for i := 0; i < 10; i++ {
		testSlice = append(testSlice, i) // память под capacity выделяется динамически, если заполнена, то увеличивается в 2 раза
		debugSlice(testSlice)

	}
	debugSlice(testSlice) */

	s := make([]int, 0, 0) // способ добавления элементов
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3, 4, 5, 6)
	debugSlice(s)

	testSlice := make([]int, 0, cap(s))
	testSlice = append(testSlice, s...) // ... это спрэд(оператор распространения),
	debugSlice(testSlice)

	for i, v := range s { // рекомендовано пробегать по слайсу рэнджем
		s[i] = v * v
	}
	debugSlice(s)
}
