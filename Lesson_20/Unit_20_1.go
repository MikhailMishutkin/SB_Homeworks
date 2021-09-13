//Вариант решения домашнего задания с дополнительным условием реверса массива

package main

import "fmt"

const (
	fSize = 4
	sSize = 7

	aSize = 11
)

func merge(A [fSize]int, B [sSize]int) [aSize]int {
	var result [aSize]int

	j, k := 0, 0

	for i := 0; i < aSize; i++ {
		if j >= len(A) {
			result[i] = B[k]
			k++
			continue
		} else if k >= len(B) {
			result[i] = A[j]
			j++
			continue
		}
		if A[j] > B[k] {
			result[i] = B[k]
			k++
		} else {
			result[i] = A[j]
			j++
		}
	}
	return result
}

func reverse(C [aSize]int) [aSize]int {
	for i := 0; i < len(C)/2; i++ {
		C[i], C[aSize-i-1] = C[aSize-i-1], C[i]

	}

	return C
}

func main() {
	a := [fSize]int{1, 3, 5, 7}
	b := [sSize]int{2, 4, 6, 8, 9, 10, 11}

	fmt.Println(a)
	fmt.Println(b)

	c := merge(a, b)
	fmt.Println(c)

	c = reverse(c)
	fmt.Println(c)
}
