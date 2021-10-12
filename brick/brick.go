package main

import (
	"fmt"
)

func main() {
	fmt.Println("This program program determines if the brick can be inserted in the hole.")
	fmt.Println("Enter the hole dimensions (2 numbers):")
	var a, b float64 
	fmt.Scanln(&a, &b)
	fmt.Println("Enter the brick dimensions (3 numbers):")
	var c, d, e float64
	fmt.Scanln(&c, &d, &e)
	if a >= d && b >= e || b >= d && a >= e || a >= c && b >= e || a >= e && b >= c || a >= c && b >= d || a >= d && b >= c {
		fmt.Println("The brick can be inserted in the hole.")
	} else {
		fmt.Println("The brick cannot be inserted in the hole.")
	}
}