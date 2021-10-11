package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program determines if triangle exists with sides lenghts that are given:")
	fmt.Println("If yes, program calculates the area of the triangle")
	var a, b, c float64
	fmt.Println("Enter three triangle sides:")
	fmt.Scan(&a, &b, &c)
	if a+b > c && a+c > b && c+b > a {
		s := (a + b + c) / 2
		fmt.Println("Triangle area:", math.Sqrt(s*(s-a)*(s-b)*(s-c)))
	} else {
		fmt.Println("A triangle with such sides doesn't exist.")
	}
	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("Lenght should be > 0")
	}
}
