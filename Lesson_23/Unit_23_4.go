// сортировка пузырьком

package main

import "fmt"

const size = 10

func bubbleSort(input [size]int64) [size]int64 {
	for i := size; i > 0; i-- {
		for j := 1; j < i; j++ {
			if input[j-1] > input[j] {
				tmpInput := input[j]
				input[j] = input[j-1]
				input[j-1] = tmpInput
			}
		}
	}
	return input
}

func main() {
	a := [size]int64{10, 44, 32, 8, 4, 5, 1, 7, 23, 29}
	fmt.Println("---UNSORTED---")
	fmt.Println(a)

	b := bubbleSort(a)
	fmt.Println("---SORTED---")
	fmt.Println(b)
}
