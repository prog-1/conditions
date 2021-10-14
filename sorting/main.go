package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program sorts three numbers.")
	var a, b, c float64
	fmt.Println("Enter three numbers:")
	fmt.Scan(&a, &b, &c)
	if a < b && a > c && b > c {
		a, b, c = c, a, b
	} else if b < a && c < a && b < c {
		a, b, c = b, c, a
	} else if a < c && c < b && a < b {
		b, c = c, b
	} else if c < b && b < a && c < a {
		a, c = c, a
	} else if b < a && a < c && b < c {
		a, b = b, a
	} else if a == b && b > c && a > c {
		a, c = c, a
	} else if a == c && a > b {
		a, b = b, a
	} else if a == c && a < b {
		b, c = c, b
	}
	fmt.Println("Numbers in a sorted order:", a, b, c)
}
