package main

import "fmt"

func main() {
	fmt.Println("The program solves a linear inequality Ax + B > 0.")
	fmt.Print("Enter A and B: ")
	var x, y float64
	fmt.Scan(&x, &y)
	if x == 0 && y == 0 {
		fmt.Println("No solutions.")
	} else if x == 0 {
		fmt.Println("No solutions.")
	} else if y == 0 {
		if x > 0 {
			fmt.Println("x > 0")
		} else {
			fmt.Println("x < 0")
		}
	} else if x < 0 {
		result := -y / x
		fmt.Println("x <", result)
	} else {
		result := -y / x
		fmt.Println("x >", result)
	}
}
