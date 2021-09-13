package main

import "fmt"

const (
	rows = 3
	cols = 5
)

func maximum(input [cols]int) int {
	max := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}
func findMaxInMatrix(matrix [rows][cols]int) [rows]int { //вычисляем максимум в каждом массиве матрицы
	var result [rows]int
	for i := 0; i < len(matrix); i++ {
		lMax := maximum(matrix[i])
		result[i] = lMax
	}
	return result
}

func sumAll(matrix [rows][cols]int) int { // сумма всех элементов матрицы
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			sum = sum + matrix[i][j]
		}
	}
	return sum
}

func sumNotEven(matrix [rows][cols]int) int { // сумма чётных массивов в
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j%2 != 0 {
				sum = sum + matrix[i][j]
			}
		}

	}
	return sum
}

func main() {
	matrix := [rows][cols]int{
		{1, 2, 3, 4, 5},
		{3, 2, 10, 100, 101},
		{40, 50, -10, -20, -30},
	}
	fmt.Println(findMaxInMatrix(matrix))
	fmt.Println(sumAll(matrix))
	fmt.Println(sumNotEven(matrix))
}
