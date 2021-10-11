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
		S := (a + b + c) / 2
		Area := math.Sqrt( S * (S - a) * (S - b) * (S - c))
		fmt.Println("Triangle area:", Area)
	} else {
		fmt.Println("A triangle with this sides does not exist.")
	}
}