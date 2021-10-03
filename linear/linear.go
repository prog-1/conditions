package main

import "fmt"

func main() {
	fmt.Println("The program solves a linear inequality Ax + B > 0.")
	fmt.Print("Enter A and B: ")
	var a, b float64
	fmt.Scan(&a, &b)
	if a == 0 && b == 0 {
		fmt.Println("No solutions.")
	} else if a == 0 {
		fmt.Println("No solutions.")
	} else if b == 0 {
		if a > 0 {
			fmt.Println("x > 0")
		} else {
			fmt.Println("x < 0")
		}
	} else if a < 0 {
		sol := -b / a //sol = solution\\
		fmt.Println("x <", sol)
	} else {
		sol := -b / a
		fmt.Println("x >", sol)
	}
}
