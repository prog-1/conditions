package main

import "fmt"

func main() {
	fmt.Println("The program solves a linear inequality Ax + B > 0.")
	var a, b float64
	fmt.Print("Enter A and B: ")
	fmt.Scan(&a, &b)
	fmt.Println("x >", -b/a)
}
