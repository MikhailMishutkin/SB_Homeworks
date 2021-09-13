package main

import "fmt"

func main() {

	var a float32 = 0

	var b float32 = 0

	for i := 0; i < 10; i++ {

		a += 0.1

	}

	for i := 0; i < 16; i++ {

		b += 0.0625

	}

	fmt.Printf("a = %v\n", a)

	fmt.Printf("b = %v\n", b)

}
