package main

import "fmt"

func main() {
	var A, B, P, Q, R float64
	fmt.Println("Enter the hole dimensions (2 numbers):")
	fmt.Scan(&A, &B)
	fmt.Println("Enter the brick dimensions (3 numbers):")
	fmt.Scan(&P, &Q, &R)
	if ((A >= R) && (B >= Q)) || ((A >= R) && (B >= P)) || ((A >= Q) && (B >= R)) || ((A >= Q) && (B >= P)) || ((A >= P) && (B >= R)) || ((A >= P) && (B >= Q)) {
		fmt.Println("The brick can be inserted in the hole!")
	} else {
		fmt.Println("The brick cannot be inserted in the hole.")
	}
}
