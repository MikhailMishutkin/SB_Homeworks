package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Введите сумму вклада")
	var deposit float64
	fmt.Scan(&deposit)

	fmt.Println("Введите срок вклада")
	var pd float64
	fmt.Scan(&pd)

	fmt.Println("Введите процент вклада")
	var pt float64
	fmt.Scan(&pt)

	rest := 0.0
	for i := 0.0; i < pd*12; i++ {
		cap := deposit * pt / 100                   // размер капитализации в итерацию
		deposit += cap                              // увеличиваем сумму депозита на капитализацию
		roundCap := deposit                         // сохраняем сумму депозита до округления
		deposit = (math.Trunc(deposit * 100)) / 100 // выделение копеек для округления
		rest = math.Abs(deposit-roundCap) + rest    // выделение остатка и суммирование
	}
	fmt.Println(deposit, rest)

}
