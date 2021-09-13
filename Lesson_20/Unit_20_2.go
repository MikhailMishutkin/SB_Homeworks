package main

import "fmt"

const (
	rows = 3
	cols = 4
)

func main() {

	matrix := [rows][cols]int{
		{1, 0, 1, 0},
		{2, 3, 3, 2},
		{4, 5, 5, 5},
	}
	//fmt.Println(matrix)
	for i := 0; i < len(matrix); i++ { //вывод матрицы как она задана
		fmt.Println(matrix[i])
	}

	for i := 0; i < len(matrix); i++ { // вывод отдельного элемента матрицы построчно
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}
}
