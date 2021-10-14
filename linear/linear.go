package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program solves a linear inequality Ax + B > 0.")
	fmt.Print("Enter A and B:")
	var a, b float64
	fmt.Scan(&a, &b)
	var x float64
	if a == 0 {
		if b < 0 {
			fmt.Println("Infinite roots")
			return
		} else {
			fmt.Println("Error! You can not divide by 0.")
			return
		}

	}
	if a > 0 {
		x = (-b / a)
		fmt.Println("X >", x)
		return
	}
	if a < 0 {
		x = (-b / a)
		fmt.Println("X <", x)
	}

}
