package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds solutions of quadratic equation.")
	var a, b ,c float64
	fmt.Println("Enter the coefficients A, B and C:")
	fmt.Scanln(&a, &b, &c)
	if a == 0 {
		if b == 0 {
			if c == 0 {
				fmt.Println(infinite roots)
			} else {
				fmt.Println(no roots)
			}
		} else {
			x :=  -c / b
			fmt.Println("x =", x)
		}
		return
	}
	D =: b * b - 4 * a * c
	if D < 0 {
		fmt.Println(no real roots)
	} else if D == 0 {
		x := (-b + math.Sqrt(D)) / (2 * a)
		fmt.Println("x1 = x2 =", x)
	} else if D > 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b + math.Sqrt(D)) / (2 * a)
		fmt.Println("x1 =", x1)
		fmt.Println("x2 =", x2)
	}
}