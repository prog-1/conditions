package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program finds solutions of a linear equation.")
	fmt.Println("Enter the coefficients A and B:")
	var a, b float64
	fmt.Scan(&a, &b)
	if a == 0 {
		if b == 0 {
			fmt.Println("infinite roots")
		} else {
			fmt.Println("no roots")
		}

	} else if a > 0 {
		{
			fmt.Println("x >", -b/a)
		}
	} else {
		fmt.Println("x <", -b/a)
	}
}
