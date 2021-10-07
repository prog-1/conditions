package main

import "fmt"

func main() {
	fmt.Println("The program solves a linear inequality Ax + B > 0.")
	fmt.Println("Enter A and B:")
	var A, B float64
	fmt.Scan(&A, &B)
	if A == 0 && B == 0 {
		fmt.Println("No decision")
	} else if A == 0 {
		fmt.Println("No decision")
	} else if B == 0 {
		if A > 0 {
			fmt.Println("x > 0")
			if A < 0 {
				fmt.Println("x < 0")
			}
		}
	}
	if A < 0 {
		fmt.Println("x <", -B/A)
	} else if A > 0 {
		fmt.Println("x >", -B/A)
	}
}
