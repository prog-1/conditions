package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the hole dimensions (2 numbers):")
	var x, y float64
	fmt.Scan(&x, &y)
	fmt.Print("Enter the brick dimensions (3 numbers):")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if x >= a && y >= b {
		fmt.Println("The brick can be inserted in the hole!")
		return
	}
	if x >= b && y >= a {
		fmt.Println("The brick can be inserted in the hole!")
		return
	}
	if x >= c && y >= a {
		fmt.Println("The brick can be inserted in the hole!")
		return
	}
	if x >= a && y >= c {
		fmt.Println("The brick can be inserted in the hole!")
		return
	}
	if x >= b && y >= c {
		fmt.Println("The brick can be inserted in the hole!")
		return
	}
	if x >= c && y >= b {
		fmt.Println("The brick can be inserted in the hole!")
		return
	}
	fmt.Println("The brick cannot be inserted in the hole.")

}
