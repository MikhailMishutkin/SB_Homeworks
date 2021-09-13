// Несколько возвращаемых значений из функций

package main

import (
	"fmt"
	"log"
)

const size = 10

func minAndMax(input [size]int) (int, int) {
	if size == 0 {
		return -1, 1
	}
	max := input[0]
	min := input[0]

	for i := 0; i < size; i++ {
		if input[i] > max {
			max = input[i]
		}
		if input[i] < min {
			min = input[i]
		}
	}
	return min, max
}

func ignoreSecond() (int, int) {
	return 3, 7
}

func calculateWithError(input [size]int) (int, error) {
	if size == 0 {
		return -1, fmt.Errorf("empty size of input")
	}
	sum := 0
	for i := 0; i < size; i++ {
		sum = sum + input[i]
	}
	return sum, nil
}

func sumAndDiff(a, b int) (sum, diff int) { // несколько именованных возвращаемых значений
	sum = a + b
	diff = a - b
	return
}

func main() {
	data := [size]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(data)

	min, max := minAndMax(data) // как получить из функции несколько значений в переменные
	fmt.Println("max is", max)
	fmt.Println("min is", min)

	a, _ := ignoreSecond() // избегаем вывод второго значения с помощью _
	fmt.Println(a)

	fmt.Println(data)
	sum, err := calculateWithError(data)
	if err != nil {
		log.Fatalf("calculate error: %v\n", err)
	}
	fmt.Println("sum is ", sum)

	s, d := sumAndDiff(10, 5)
	fmt.Println("sum is ", s, "diff is", d)
}
