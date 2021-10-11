package main

import "fmt"

func main() {
	fmt.Println("The program solves a linear inequality Ax + B > 0.")
	fmt.Println("Enter A and B:")
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println("x>", -b/a)
}
