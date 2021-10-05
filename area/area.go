package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter three triangle sides:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if a+b > c && a+c > b && b+c > a {
		p := (a + b + c) / 2
		area := math.Sqrt(p * (p - a) * (p - b) * (p - c))
		fmt.Printf("Triangle area: %v", area)
	} else {
		fmt.Println("A triangle with such sides doesn't exist.")
	}
}
