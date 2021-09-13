package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("введите до какого знака важна точность значения функции")
	var epsilon float64
	fmt.Scan(&epsilon)

	epsilon1 := 1 / math.Pow(10, epsilon) //вычисляем погрешность
	fmt.Println("введите х для которого необходимо вычислить значение функции")
	var x float64
	fmt.Scan(&x)

	result := 1.0
	prevResult := 0.0
	fact := 1.0
	n := 1.0

	for {
		if n > 1 {
			fact *= n
		}
		result += (math.Pow(x, n) / float64(fact))
		if math.Abs(result-prevResult) < epsilon1 {
			fmt.Println(result)
			break
		}
		n++
		prevResult = result
		fmt.Println(result)
	}
}
