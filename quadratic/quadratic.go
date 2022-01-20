package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds solutions of quadratic equation.")
	fmt.Print("Enter the coefficients A, B and C: ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	d := b*b - 4*a*c
	if a == 0 {
		fmt.Println("x =", -c/b)
	} else if d < 0 {
		fmt.Println("There are no solutions")
	} else {
		fmt.Println("x1 =", (-b+math.Sqrt(d))/2)
		fmt.Println("x2 =", (-b-math.Sqrt(d))/2)
	}
}
