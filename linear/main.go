package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program solves a linear inequality Ax + B > 0.")
	var A, B float64
	fmt.Println("Enter A and B:")
	fmt.Scanln(&A, &B)
	if A < 0 {
		x := -B / A
		fmt.Println("x <", x)
	} else if A > 0 {
		x := -B / A
		fmt.Println("x >", x)
	} else if A == 0 {
		fmt.Println("ERROR: cannot divide by zero")
	}
}
