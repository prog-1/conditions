package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds solutions of quadratic equation.")
	var a, b, c float64
	fmt.Println("Enter the coefficients A, B and C:")
	fmt.Scan(&a, &b, &c)
	if a == 0 {
		fmt.Println("x=", -c/b)
		if b == 0 {
			fmt.Println("No roots")
			if c == 0 {
				fmt.Println("Infinite roots")
			}
		}
	} else {
		D := b*b - 4*a*c
		if D < 0 {
			fmt.Println("No roots")
			if D == 0 {
				fmt.Println("x1, x2 =", -b/(2*a))
				if D > 0 {
					fmt.Println("x1=", (-b+math.Sqrt(D))/(2*a))
					fmt.Println("x2=", (-b-math.Sqrt(D))/(2*a))
				}
			}
		}
	}
}
