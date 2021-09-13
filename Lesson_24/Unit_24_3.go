// Добавление элементов  в подсрез, удаление элементов слайса

package main

import "fmt"

func debugSlice(input []int) {
	fmt.Printf("data: %+v\n", input)
	fmt.Printf("len: %+v\n", len(input))
	fmt.Printf("cap: %+v\n", cap(input))
}

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func anotherRemove(slice []int, index int) []int {
	slice[len(slice)-1], slice[index] = slice[index], slice[len(slice)-1] // элемент, который надо удалить поставлен в конец массива
	return slice[:len(slice)-1]                                           // возвразаем обрезав последний элемент
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	debugSlice(s)

	d := s[1:3] // подсрез, граница среза = крайний индекс минус 1 (т.е. в подсрезе будет индекс 1 и 2)
	debugSlice(d)

	e := s[3:5] // capacity среза равна capacity массива минус первый индекс среза (в данном случае 6-3)
	debugSlice(e)

	// copy()
	d1 := make([]int, 2, 2)
	copy(d1, s[2:4]) // два аргумента - в какой массива из какого
	fmt.Println("D1", d1)

	//удаление через append - ёмкость не меняется
	s = remove(s, 2)
	debugSlice(s)

	// второй вариант удаления
	s = anotherRemove(s, 3)
	debugSlice(s)
}
