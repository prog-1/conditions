package main

import "fmt"

func main() {
	fmt.Println("The program determines whether a brick can be inserted into the hole or not.")
	fmt.Println("Enter the hole dimensions (2 numbers):")
	var a, b float64
	fmt.Scanln(&a, &b)
	fmt.Println("Enter the brick dimensions (3 numbers):")
	var p, q, r float64
	fmt.Scanln(&p, &q, &r)
	if a >= p && (b >= q || b >= r) ||
		a >= q && (b >= p || b >= r) ||
		a >= r && (b >= p || b >= q) {
		fmt.Println("The brick can be inserted in the hole!")
	} else {
		fmt.Println("The brick cannot be inserted in the hole.")
	}
}
