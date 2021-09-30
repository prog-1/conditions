package main

import "fmt"

func main() {
	fmt.Println("This program determines if the brick can be inserted in the hole.")
	fmt.Print("Enter the hole dimensions (2 numbers): ")
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Print("Enter the brick dimensions (3 numbers): ")
	var p, q, r float64
	fmt.Scan(&p, &q, &r)
	if a >= q && b >= r || b >= r && a >= q {
		fmt.Println("The brick can be inserted in the hole!")
	} else {
		fmt.Println("The brick cannot be inserted in the hole.")
	}
}
