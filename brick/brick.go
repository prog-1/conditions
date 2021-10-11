package main

import "fmt"

func main() {
	fmt.Println("This program determines if the brick can be inserted in the hole.")
	fmt.Print("Enter the hole dimensions (2 numbers): ")
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Print("Enter the brick dimensions (3 numbers): ")
	var x, y, z float64
	fmt.Scan(&x, &y, &z)
	if a >= y && b >= z || b >= y && a >= z || a >= x && b >= z || a >= z && b >= x || a >= x && b >= y || a >= y && b >= x {
		fmt.Println("The brick can be inserted in the hole!")
	} else {
		fmt.Println("The brick cannot be inserted in the hole.")
	}
}
