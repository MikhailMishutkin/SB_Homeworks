package main

import (
	"fmt"
)

const ()

var (
	a uint8  = 0
	b uint16 = 0
	c uint32 = 0
)

var countA, countB int

func main() {
	for {
		a++
		b++
		c++
		if a == 0 {
			countA++
		}
		if b == 0 {
			countB++
		}
		if c == 0 {
			break
		}
		// fmt.Println(a, b)
	}

	fmt.Println(countA, countB)
}
