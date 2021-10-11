package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program that determines if the brick can be inserted in the hole.")
	var a, b float64
	fmt.Println("Enter the hole dimensions (2 numbers):")
	fmt.Scan(&a, &b)
	var p, q, r float64
	fmt.Println("Enter the brick dimensions (3 numbers):")
	fmt.Scan(&p, &q, &r)
	if a >= p && b >= q || a >= q && b >= p || a >= r && b >= q || a >= q && b >= r || a >= r && b >= p || a >= p && b >= r {
		fmt.Println("The brick can be inserted in the hole")
	} else {
		fmt.Println("The brick cannot be inserted in the hole")
	}
	if a <= 0 || b <= 0 || p <= 0 || q <= 0 || r <= 0 {
		fmt.Println("Lenght should be > 0")
	}
}
