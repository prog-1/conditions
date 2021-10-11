package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds solutions of quadratic equation.")
	fmt.Println("Enter the coefficients A, B and C:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	var d = b*b - 4*a*c
	if a == 0 {
		var x = -b / c
		fmt.Println("x=", x)
	} else if d > 0 {
		var x1 = (-b + math.Sqrt(d)) / 2 * a
		var x2 = (-b - math.Sqrt(d)) / 2 * a
		fmt.Println("x1:", x1, "x2:", x2)
	} else if d == 0 {
		fmt.Println("x:", -b/2*a)
	} else {
		fmt.Println("no solution")
	}

}
