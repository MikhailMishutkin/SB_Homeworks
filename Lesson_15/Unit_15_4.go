//Организация массивов. Вывод четных элементов массива

package main

import "fmt"

func main() {
	var arr [100]int
	var n int
	fmt.Println("Сколько чисел вы хотите ввести: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("arr[%v]=", i)
		fmt.Scan(&arr[i])
	}
	fmt.Println("Вывод")
	for i := 0; i < (n+1)/2; i++ {
		index := i*2 + 1
		fmt.Printf("arr[%v]=%v\n", index, arr[index])
	}
}
