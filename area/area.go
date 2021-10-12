package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter three triangle sides:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	var s = (a + b + c) / 2
	var A = math.Sqrt((s * (s - a) * (s - b) * (s - )))
	if A > 0 {
		fmt.Prinln("TRiangle area:" , A)
	} else {
		fmt.Println("Triangle with such sides doesn't exist.")
	}
}