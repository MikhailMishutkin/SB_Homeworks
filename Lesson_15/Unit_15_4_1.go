//Организация массивов. Расчёт скидки

package main

import "fmt"

func main() {
	var price [100]float64
	var n int
	fmt.Println("Сколько товаров вы хотите ввести: ")
	fmt.Scan(&n)
	sum := 0.0
	for i := 0; i < n; i++ {
		fmt.Printf("Введите стоимость товара %v ", i+1)
		fmt.Scan(&price[i])
		sum += price[i]
	}
	endsum := 0.0
	for i := 0; i < n; i++ {
		if sum > 5000 {
			price[i] *= 0.95
		}
		endsum += price[i]
		fmt.Printf("%v товар стоит %v р. \n", i+1, price[i])

	}
	fmt.Printf("Итого %v р. (скидка %v)", endsum, endsum-sum)
}
