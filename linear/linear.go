package main

import "fmt"

func main() {
	fmt.Println("The program solves a linear inequality Ax + B > 0.")
	fmt.Println("Enter A and B:")
	var a, b float64
	fmt.Scanln(&a, &b)
	result := -b / a
	if a > 0 && b != 0 {
		fmt.Println("x >", result)
	} else if a < 0 && b != 0 {
		fmt.Println("x <", result)
	} else if a > 0 && b == 0 {
		fmt.Println("x > 0")
	} else if a < 0 && b == 0 {
		fmt.Println("x < 0")
	} else if a == 0 && b != 0 {
		fmt.Println("x ∈ R")
	} else if a == 0 && b == 0 {
		fmt.Println("x ∈ ∅")
	}
}
