package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Enter three triangle sides: ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	s := ((a + b + c) / 2)
	x := math.Sqrt(s * (s - a) * (s - b) * (s - c))
	if x != math.Sqrt(s*(s-a)*(s-b)*(s-c)) {
		fmt.Println("A triangle with such sides doesn't exist.")
	} else {
		fmt.Println("Triangle area: ", x)
	}
}
