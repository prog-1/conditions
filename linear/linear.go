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
	x = (-b / a)
	fmt.Println("X >", x)

}
