package main

import "fmt"

func main() {
	fmt.Println("The program solves a linear inequality Ax + B > 0")
	var a, b float64
	fmt.Println("Enter A and B:")
	fmt.Scan(&a, &b)
	if a == 0 && b == 0 {
		fmt.Println("No decision")
	} else if a == 0 {
		fmt.Println("No decision")
	} else if b == 0 {
		if a > 0 {
			fmt.Println("x > 0")
		} else if a < 0 {
			fmt.Println("x < 0")
		}
		if a < 0 {
			fmt.Println("x <", -b/a)
			if a > 0 {
				fmt.Println("x >", -b/a)
			}
		}
	}
}
