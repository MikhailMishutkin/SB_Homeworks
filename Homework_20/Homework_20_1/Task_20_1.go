/* Задание 1. Подсчёт определителя
Что нужно сделать
Напишите функцию, вычисляющую определитель матрицы размером 3 × 3. */

package main

import "fmt"

const (
	rows  = 3
	cols  = 3
	fCols = 5
)

func methSar(A [rows][cols]int, B [rows][cols]int) [rows][fCols]int { // добавление четвертой и пятой колонки по методу Саррюса
	var uniteMatrixSar [rows][fCols]int
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			uniteMatrixSar[i][j] = A[i][j]
		}
	}
	for i := 0; i < len(B); i++ {
		for j := 1; j < len(B[0]); j++ {
			uniteMatrixSar[i][len(A)+j-1] = B[i][j-1]
		}
	}
	return uniteMatrixSar
}

func diagDifference(A [rows][fCols]int) int { // определитель - разница произведений элементов матрицы по диагоналям
	prod, prod1, prod2, bProd, bProd1, bProd2 := 1, 1, 1, 1, 1, 1 // переменные под результаты вычисления произведения диагоналей в обе стороны
	j, k, l := 0, 0, 0                                            // счётчики для обратного хода
	for i := 0; i < len(A); i++ {
		prod = prod * A[i][i]
	}
	for i := 0; i < len(A); i++ {
		prod1 = prod1 * A[i][i+1]
	}
	for i := 0; i < len(A); i++ {
		prod2 = prod2 * A[i][i+2]
	}

	for i := len(A); i > 0; i-- {
		bProd = bProd * A[j][fCols-j-1]
		j++
	}
	for i := len(A); i > 0; i-- {
		bProd1 = bProd1 * A[k][fCols-k-2]
		k++
	}
	for i := len(A); i > 0; i-- {
		bProd2 = bProd2 * A[l][fCols-l-3]
		l++
	}

	difference := prod + prod1 + prod2 - bProd - bProd1 - bProd2 // формула из вики

	return difference

}

func main() {
	matrix := [rows][cols]int{
		{1, 2, 3},
		{10, 8, 16},
		{18, 4, 9},
	}

	//fmt.Println(methSar(matrix, matrix))
	e := methSar(matrix, matrix)
	for i := 0; i < len(matrix); i++ { //вывод матрицы как она задана расширенный вид
		fmt.Println(e[i])
	}
	fmt.Println(diagDifference(e))
	//fmt.Println(len(e))
	//fmt.Println(difference)
}
