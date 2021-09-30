package main

import (
	"fmt"
	"math"
)

func max(a, b float64) float64 {
	return (a + b + math.Abs(a-b)) / 2
}
func main() {
	var a, b, c float64
	fmt.Println("This program calculates area of triangle")
	fmt.Print("Enter three triangle sides: ")
	fmt.Scan(&a, &b, &c)

	if a+b < c {
		fmt.Println("Triangle does not exist")
		return
	}
	if a+c < b {
		fmt.Println("Triangle does not exist")
		return
	}
	if b+c < a {
		fmt.Println("Triangle does not exist")
		return
	}
	var p, S float64
	p = (a + b + c) / 2.0
	S = math.Sqrt(p * (p - a) * (p - b) * (p - c))
	fmt.Println("Triangle area:", S)

}
