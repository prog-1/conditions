package main

import "fmt"

func main() {
	fmt.Println("The program sorts three numbers.")
	fmt.Println("Enter the numbers:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	if a > b && a > c {
		if b > c {
			a, b, c = c, b, a
		} else {
			a, b, c = b, c, a
		}
	} else if b < c && a < c {
		if a > b {
			a, b, c = b, a, c
		} else {
			a, b, c = a, b, c
		}
	} else if a < b && b > c {
		if a > c {
			a, b, c = c, a, b
		} else {
			a, b, c = a, c, b
		}
	} else {
		fmt.Println("Numbers are equal.")
		return
	}

	fmt.Println("Numbers in a sorted order:", a, b, c)
}
