package main

import "fmt"

func main() {
	fmt.Println("The program solves a linear inequality Ax + B > 0.")
	fmt.Println("Enter A and B:")
	var a, b float64
	fmt.Scanln(&a, &b)
	if a == 0 && b != 0 { // To avoid division by 0.
		fmt.Println("x ∈ ∅")
		return
	} else if a == 0 && b == 0 {
		fmt.Println("x ∈ R")
		return
	}
	result := -b / a
	if a > 0 && b != 0 || a > 0 && b == 0 { // Or, to avoid division by 0, result := -b / a could be placed in these 2 if.
		fmt.Println("x >", result)
	} else if a < 0 && b != 0 || a < 0 && b == 0 {
		fmt.Println("x <", result)
	}
}
