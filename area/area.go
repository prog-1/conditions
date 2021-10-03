package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("This program calculates areaof the triangle.")
	fmt.Println("Enter three triangle sides: ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if a+b > c && b+c > a && c+a > b {
		s := (a + b + c) / 2
		area := math.Sqrt(s * (s - a) * (s - b) * (s - c))
		fmt.Println("Triangle area:", area)
	} else {
		fmt.Println("A triangle with such sides doesn't exist.")
	}
}
