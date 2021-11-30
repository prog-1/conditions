package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program calculates the area of the triangle.")
	fmt.Println("Enter three triangle sides:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	if a+b > c && a+c > b && b+c > a {
		s := (a + b + c) / 2
		A := math.Sqrt(s * (s - a) * (s - b) * (s - c))
		fmt.Println("Triangle area:", A)
	} else {
		fmt.Println("A triangle with such sides doesn't exist.")
	}
}
