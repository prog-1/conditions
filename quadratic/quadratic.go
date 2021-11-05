package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program finds solutions of quadratic equation.")
	fmt.Println("Enter the coefficients A, B and C:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	// d := (b * b) - 4*a*c
	// if d < 0 {
	// 	fmt.Println("No roots")
	// 	if d == 0 {
	// 		fmt.Println("x1=x2=", (2*a)/-b)
	// 		if d == 0 {
	// 			fmt.Println("x1=", (-b+math.Sqrt(d))/(a*2))
	// 			fmt.Println("x2=", (-b-math.Sqrt(d))/(a*2))
	// 			if a == 0 {
	// 				fmt.Println("x=", -c/b)
	// 				if b == 0 {
	// 					fmt.Println("No roots")
	// 					if c == 0 {
	// 						fmt.Println("Infinitive roots")
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}
