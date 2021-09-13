package main

import (
	"fmt"
	"math"
)

func main() {
	var a int

	a = math.MaxInt8
	fmt.Println(a)
	b := float64(a)
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)

	a = math.MinInt8
	fmt.Println(a)
	b = float64(a)
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)

	a = math.MaxUint8
	fmt.Println(a)
	b = float64(a)
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)

	a = math.MaxInt16
	fmt.Println(a)
	b = float64(a)
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)

	a = math.MinInt16
	fmt.Println(a)
	b = float64(a)
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)

	a = math.MaxUint16
	fmt.Println(a)
	b = float64(a)
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)

	a = math.MaxInt32
	fmt.Println(a)
	b = float64(a)
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)

	a = math.MinInt32
	fmt.Println(a)
	b = float64(a)
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)

	a = math.MaxUint32
	fmt.Println(a)
	b = float64(a)
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)

	a = math.MaxInt64
	fmt.Println(a)
	b = float64(a)
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)

	a = math.MinInt64
	fmt.Println(a)
	b = float64(a)
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)

	//	a = math.MaxUint64 - 1
	//	fmt.Println(a)
	//	b = float64(a)
	//	fmt.Println(b)
	//	a = int(b)
	//	fmt.Println(a)
}
