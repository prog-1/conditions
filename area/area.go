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
	var A = math.Sqrt((s * (s - a) * (s - b) * (s - c)))
	if A > 0 {
		fmt.Println("Triangle area:", A)
	} else {
		fmt.Println("A triangle with such sides doesn't exist.")
	}
}
