// Разбор домашки 20
// Вычисление определителя матрицы
// Умножение матриц

package main

import "fmt"

const (
	rows  = 3
	cols  = 3
	oRows = 3
	oCols = 5
	nCols = 4
)

func determinant(matrix [rows][cols]int64) int64 {
	x := matrix[1][1]*matrix[2][2] - matrix[2][1]*matrix[1][2]
	y := matrix[1][0]*matrix[2][2] - matrix[2][0]*matrix[1][2]
	z := matrix[1][0]*matrix[2][1] - matrix[2][0]*matrix[1][1]

	determinant := matrix[0][0]*x - matrix[0][1]*y + matrix[0][2]*z
	return determinant
}

func multiply(m1 [oRows][oCols]int64, m2 [oCols][nCols]int64) [oRows][nCols]int64 {
	var m [oRows][nCols]int64
	for i := 0; i < oRows; i++ {
		for j := 0; j < nCols; j++ {
			for k := 0; k < oCols; k++ {
				m[i][j] = m[i][j] + m1[i][k]*m2[k][j]
			}

		}

	}
	return m
}

func main() {
	mm := [rows][cols]int64{
		{1, 0, 10},
		{10, 11, 12},
		{10, 11, 13},
	}
	fmt.Println(determinant(mm))

	m1 := [oRows][oCols]int64{
		{1, 2, 3, 4, 5},
		{2, 4, 6, 8, 9},
		{9, 7, 3, 5, 1},
	}
	m2 := [oCols][nCols]int64{
		{2, 3, 4, 5},
		{9, 8, 7, 6},
		{4, 6, 8, 1},
		{1, 9, 3, 7},
		{8, 2, 6, 4},
	}
	fmt.Println(multiply(m1, m2))
}
