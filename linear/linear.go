package main

import "fmt"

func main() {
	fmt.Println("The program solves a linear inequality Ax + B > 0.")
	fmt.Println("Enter A and B:")
	var A, B float64
	fmt.Scan(&A, &B)
	if A > 0 {
		x := -B / A
		fmt.Printf("x > %v", x)
	} else if A < 0 {
		x1 := -B / A
		fmt.Printf("x < %v", x1)
	} else {
		if B <= 0 {
			fmt.Println("No answer.")
		} else {
			fmt.Println("Answer is all possible numbers.")
		}
	}

}
