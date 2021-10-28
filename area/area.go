package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program determines if a triangle with such side lengths exists. If yes, calculates the area of the triangle")
	var a, b, c float64
	fmt.Println("Enter three triangle sides:")
	fmt.Scan(&a, &b, &c)
	if a+b > c && a+c > b && b+c > a {
		semip := (a + b + c) / 2
		area := math.Sqrt(semip * (semip - a) * (semip - b) * (semip - c))
		fmt.Println("Triangle area:", area)
	} else {
		fmt.Println("A triangle with such sides doesn't exist.")
	}
	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("Lenght need to be greater than 0")
		return
	}
}
