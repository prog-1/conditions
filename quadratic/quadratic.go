package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds solutions of a quadratic equation.")
	fmt.Println("Enter the coefficients A, B and C:")
	var A, B, C float64
	fmt.Scan(&A, &B, &C)
	if A == 0 {
		if B != 0 {
			x := -C / B
			fmt.Printf("x = %v", x)
		} else {
			if C != 0 {
				fmt.Println("No real answers.")
			} else {
				fmt.Println("Answer is all possible numbers.")
			}
		}
	} else {
		var D float64 = B*B - 4*A*C
		if D > 0 {
			x1 := (-B + math.Sqrt(D)) / (2 * A)
			x2 := (-B - math.Sqrt(D)) / (2 * A)
			fmt.Printf("x1 = %v\n", x1)
			fmt.Printf("x2 = %v", x2)
		} else if D < 0 {
			fmt.Println("No answer.")
		} else {
			x := (-B + math.Sqrt(D)) / (2 * A)
			fmt.Printf("x1 = x2 = %v", x)
		}

	}
}
