package main

import "fmt"

const (
	rows = 3
	cols = 4
)

func sumMatrix(A [rows][cols]int, B [rows][cols]int) [rows][cols]int { // складываем матрицы
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			A[i][j] = A[i][j] + B[i][j]
		}
	}
	return A
}

func diagonalSum(A [rows][cols]int) int { //сумма элементов матрицы по диагонали
	sum := 0
	for i := 0; i < len(A); i++ {
		sum = sum + A[i][i]
	}
	return sum
}

func printInOneCycle(m [rows][cols]int) { // метод как вывести на печать все элементы матрицы использую только один цикл
	for i := 0; i < rows*cols; i++ {
		row := i / cols
		col := i % cols
		fmt.Println(m[row][col])
	}

}

func transpose(A [rows][cols]int) [cols][rows]int { // транспонирование матрицы (переворот колонок в ряды)
	var transposed [cols][rows]int
	for i := 0; i < len(A[0]); i++ {
		for j := 0; j < len(A); j++ {
			transposed[i][j] = A[j][i]
		}
	}
	return transposed
}

func main() {
	matrix := [rows][cols]int{
		{10, 10, 10, 10},
		{10, 20, 10, 20},
		{-10, -20, -10, -10},
	}
	fmt.Println(sumMatrix(matrix, matrix))
	fmt.Println(diagonalSum(matrix))
	fmt.Println(transpose(matrix))
	printInOneCycle(matrix)
	fmt.Println(len(matrix[0]))
}
