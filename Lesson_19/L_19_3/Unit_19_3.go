//Алгоритмическое решение задач с помощью массива

package main

import (
	"fmt"

	"github.com/google/gxui/math"
)

const size1 = 10

/* func twoSmallestNumbers(input [size1]int) (int, int) { // находим два минимальных значения в массиве
	min := input[0]

	for i := 0; i < len(input); i++ {
		if input[i] < min { // пробегаем по массиву и присваиваем первый минимум
			min = input[i]
		}
	}

	secondMin := math.MaxInt // присваиваем минимум максимальному инту
	for i := 0; i < len(input); i++ {
		if input[i] < secondMin && input[i] > min { // меньша максимального инта и больше min
			secondMin = input[i]
		}
	}

	return min, secondMin // работает только если элементы не задваиваются
} */

func findTwoLargestNumbers(input [size1]int) (int, int) {
	max := -math.MaxInt - 1
	secondMax := math.MaxInt - 1

	for i := 0; i < len(input); i++ {
		if input[i] > max {
			secondMax = max
			max = input[i] // переопредление глобального макс
		} else if input[i] >= secondMax { // больше max, но меньше secondMax
			secondMax = input[i]
		}
	}
	return max, secondMax
}

func main() {

	numbers := [size1]int{10, 20, 30, 40, 50, 90, 80, 70, 2, 2}

	//	fmt.Println(twoSmallestNumbers(numbers))
	fmt.Println(NewTwoSmallestNunbers(numbers))
	fmt.Println(findTwoLargestNumbers(numbers))
}
func NewTwoSmallestNunbers(input [size1]int) (int, int) {
	min := math.MaxInt
	secondMin := math.MaxInt

	for i := 0; i < len(input); i++ {
		if input[i] < min {
			secondMin = min
			min = input[i] // переопредление глобального минимума
		} else if input[i] <= secondMin { // больше min, но меньше secondMin
			secondMin = input[i]
		}
	}
	return min, secondMin
}
