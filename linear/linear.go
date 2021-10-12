package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program solves a linear inequality Ax + B > 0.")
	fmt.Println("Enter A and B:")
	var A, B float64
	fmt.Scanln(&A, &B)
	if A == 0 && B == 0 {
		fmt.Println("No decision")
	} else if A == 0 {
		fmt.Println("No decision") 
	} else if B == 0 {
		if A > 0 {
			fmt.Println("x > 0")
		} else {
			fmt.Println("x < 0")
		}
	} else if A < 0 {
		result := -B / A
		fmt.Println("x <", result)
	} else {
		result := -B / A 
		fmt.Println("x >", result)
	}
}
