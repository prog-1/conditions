package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Println("Enter the hole dimensions (2 numbers):")
	fmt.Scan(&a, &b)
	var p, q, r float64
	fmt.Println("Enter the brick dimensions (3 numbers):")
	fmt.Scan(&p, &q, &r)
	if a >= p && b >= q || b >= p && a >= q || a >= q && b >= r || a >= r && b >= q || a >= p && b >= r || a >= r && b >= p {
		fmt.Println("The brick can be inserted in the hole!")
	} else if p <= 0 || q <= 0 || r <= 0 || a <= 0 || b <= 0 {
		fmt.Println("Dimensions does not exist.")
	} else {
		fmt.Println("The brick cannot be inserted in the hole.")
	}
}
