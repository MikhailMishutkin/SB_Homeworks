// QuickSort алгоритм

package main

import "fmt"

const size = 10

func sort(input [size]int64) [size]int64 {
	for i := 0; i < size; i++ {
		var minIdx = i
		for j := i; j < size; j++ { // цикл начинается со значения i
			if input[j] < input[minIdx] {
				minIdx = j
			}
		}
		input[i], input[minIdx] = input[minIdx], input[i]
	}
	return input
}

func main() {
	a := [size]int64{10, 39, 2, 3, 10, 99, 5, 45, 54, 31}
	fmt.Println("---UNSORTED---")
	fmt.Println(a)

	b := sort(a)
	fmt.Println("---SORTED---")
	fmt.Println(b)
}
