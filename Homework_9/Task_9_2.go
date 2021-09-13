package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b int16

	fmt.Println("Введите первое число")
	fmt.Scan(&a)
	fmt.Println("Введите второе число")
	fmt.Scan(&b)

	result := int32(a) * int32(b)

	if result <= math.MinInt16 {
		fmt.Println("можно сохранить в int32")
	} else if result <= math.MinInt8 {
		fmt.Println("можно сохранить в int16")
	} else if result <= 0 {
		fmt.Println("можно сохранить в int8")
	} else if result < math.MaxUint8 {
		fmt.Println("можно сохранить в Uint8")
	} else if result < math.MaxUint16 {
		fmt.Println("можно сохранить в Uint16")
	} else {
		fmt.Println("можно сохранить в uint32")
	}

}

//&& result > math.MinInt32
// && result > math.MinInt16
// result > math.MinInt8 &&
// result > 0 &&
// result >= math.MaxUint8 &&
