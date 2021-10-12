package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter three triangle sides:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	if a+b > c && b+c > a && c+a > b {
		s := (a + b + c) / 2
		area := math.Sqrt( s * (s - a) * (s - b) * (s - c))
		fmt.Println("Triangle area:", Area)
	} else {
		fmt.Println("A triangle with this sides does not exist.")
	}
}