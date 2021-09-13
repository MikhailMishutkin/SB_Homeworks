package main

import (
	"fmt"
)

func main() {
	fmt.Println("Зеркальные билетики")
	fmt.Println("===================")

	const min = 100000
	const max = 999999

	var counter int

	for ticket := min; ticket <= max; ticket++ {
		var a, rev int
		a = ticket
		for a > 999 {
			rev = (rev * 10) + a%10
			a /= 10
		}
		firstDigs := ticket / 1000
		if rev == firstDigs {
			counter++
		}
	}
	fmt.Println("Количество зеркальных билетов", counter)
}
