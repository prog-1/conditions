package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds solutions of quadratic equation.")
	fmt.Print("Enter the coefficients A, B and C: ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if a == 0 {
		if b == 0 {
			if c == 0 {
				fmt.Println("Infinite roots")
			} else {
				fmt.Println("No real roots")
			}
		} else {
			x := -c / b
			fmt.Println("x =", x)
		}
		return
	}
	d := b*b - 4*a*c
	if d < 0 {
		fmt.Println("No real roots")
	} else if d == 0 {
		x := (-b + math.Sqrt(d)) / (2 * a)
		fmt.Println("x1 = x2 =", x)
	} else {
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Println("x1 =", x1)
		fmt.Println("x2 =", x2)
	}
}
