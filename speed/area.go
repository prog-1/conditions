package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("Enter three triangle sides:")
	fmt.Scan(&a, &b, &c)
	if a+b > c && a+c > b && c+b > a {
		s := (a + b + c) / 2
		A := math.Sqrt(s * (s - a) * (s - b) * (s - c))
		fmt.Println("Triangle area:", A)
	} else {
		fmt.Println("A triangle with such sides doesn't exist.")
	}
}
