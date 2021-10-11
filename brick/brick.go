package main

import "fmt"

func main() {
	var a, b, c, d, e float64
	fmt.Println("Enter the hole dimensions (2 numbers):")
	fmt.Scan(&a, &b)
	fmt.Println("Enter the brick dimensions (3 numbers):")
	fmt.Scan(&c, &d, &e)
	if (a >= e && b >= d) || (a >= e && b >= c) || (a >= d && b >= e) || (a >= d && b >= c) || (a >= c && b >= e) || (a >= c && b >= d) {
		fmt.Println("The brick can fit in the hole!")
	} else {
		fmt.Println("The brick can't fit in the hole.")
	}
}
