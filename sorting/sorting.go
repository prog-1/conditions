package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program sorts three numbers.")
	fmt.Println("Enter the numbers:")
	var a, b, c, scnd float64
	fmt.Scan(&a, &b, &c)
	min1 := (a + b - math.Abs(a-b)) / 2
	min := (min1 + c - math.Abs(min1-c)) / 2
	max1 := (a + b + math.Abs(a-b)) / 2
	max := (max1 + c + math.Abs(max1-c)) / 2
	if min < scnd && scnd < max {
		fmt.Printf("Numbers in a sorted order: %v %v %v", min, scnd, max)

	}
}
