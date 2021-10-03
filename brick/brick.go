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
	if a <= 0 || b <= 0 || p <= 0 || q <= 0 || r <= 0 {
		fmt.Println("ERROR: enter number greater than zero!")
		return
	}
	if a >= q && b >= r || b >= q && a >= r || a >= p && b >= r || a >= r && b >= p || a >= p && b >= q || a >= q && b >= p {
		fmt.Println("The brick can be inserted in the hole!")
	} else {
		fmt.Println("The brick cannot be inserted in the hole.")
	}
}
