package main

import "fmt"

func main() {
	fmt.Println("This program determines if the brick can be inserted in the hole.")
	fmt.Print("Enter the hole dimensions (2 numbers): ")
	var x, y float64
	fmt.Scan(&x, &y)
	fmt.Print("Enter the brick dimensions (3 numbers): ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if a > b && a > c {
		if b*c <= x*y {
			fmt.Println("The brick can be inserted in the hole!")
		} else {
			fmt.Println("The brick cannot be inserted in the hole.")
		}
	} else if b > a && b > c {
		if a*c <= x*y {
			fmt.Println("The brick can be inserted in the hole!")
		} else {
			fmt.Println("The brick cannot be inserted in the hole.")
		}
	} else if c > a && c > b {
		if a*b <= x*y {
			fmt.Println("The brick can be inserted in the hole!")
		} else {
			fmt.Println("The brick cannot be inserted in the hole.")
		}
	} else {
		fmt.Println("The brick cannot be inserted in the hole.")
	}
}
