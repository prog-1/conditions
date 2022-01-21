package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if the brick can be inserted in the hole.")
	var A, B float64
	fmt.Println("Enter the hole dimensions (2 numbers):")
	fmt.Scan(&A, &B)
	var P, Q, R float64
	fmt.Println("Enter the brick dimensions (3 numbers):")
	fmt.Scan(&P, &Q, &R)
	if A >= P && B >= Q || A >= P && B >= R || A >= Q && B >= P || A >= R && B >= P || A >= R && B >= Q || A >= Q && B >= R {
		fmt.Println("The brick can be inserted in the hole!")
	} else {
		fmt.Println("The brick cannot be inserted in the hole.")
	}
	if A <= 0 || B <= 0 || P <= 0 || Q <= 0 || R <= 0 {
		fmt.Println("There is no solution")
		return
	}
}
