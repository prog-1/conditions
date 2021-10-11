package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter three triangle sides:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if a+b > c && c+b > a && a+c > b {

		p := (a + b + c) / 2
		s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
		fmt.Print("S=", s)

	} else {
		fmt.Print("A triangle with such sides doesn't exist.")
	}

}
