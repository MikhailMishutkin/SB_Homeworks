// неограниченное число аргументов в функции

package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func generalSum(s ...int) int { // синтаксис неограниченного количества аргументов
	ss := 0
	for _, v := range s { // складываем все аргументы
		ss = ss + v
	}
	return ss
}

func conditionalCalculate(op string, x ...int) int { // вычисление в зависимости от условий с неограниченным количеством аргументов и строкой
	var result int
	if op == "+" {
		for _, v := range x {
			result = result + v
		}
	}
	if op == "*" {
		result = 1
		for _, v := range x {
			result = result * v
		}
	}
	return result
}

func anyObjects(a ...interface{}) { // использование интерфейса для вывода неограниченного количества аргументов
	fmt.Printf("%+v\n", a)
}

func main() {
	fmt.Println(generalSum(2, 3))
	fmt.Println(generalSum(2, 3, 4, 5))
	fmt.Println(generalSum(2, 3, 4, 5, 10, 10, 10, 10, 10, 10, 10))

	fmt.Println(conditionalCalculate("+", 1, 2, 3, 4))
	fmt.Println(conditionalCalculate("*", 1, 2, 3, 4))

	anyObjects(1, "222", [2]int{1, 2})
}
