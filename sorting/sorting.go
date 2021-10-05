package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program sorts three numbers.")
	fmt.Println("Enter the numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	max := math.Max(math.Max(a, b), c)
	min := math.Min(math.Min(a, b), c)
	scnd := a + b + c - (max + min)
	fmt.Printf("Numbers in a sorted order: %v %v %v", min, scnd, max)

}
