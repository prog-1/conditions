package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds solutions of a quadratic equation.")
	fmt.Println("Enter the coefficients A, B and C:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	if a == 0 {
		if b == 0 {
			if c == 0 {
				fmt.Println("infinite roots")
			} else {
				fmt.Println("no roots")
			}

		} else {
			fmt.Println("x=", -c/b)
		}

	} else {
		D := b*b - 4*a*c
		if D >= 0 {
			if D > 0 {
				fmt.Println("x1=", ((-b)+math.Sqrt(D))/(2*a))
				fmt.Println("x2=", ((-b)-math.Sqrt(D))/(2*a))
				return
			} else {

				fmt.Println("x=", (-b)/(2*a))

			}
		} else {
			fmt.Println("no roots")
		}
	}
}
