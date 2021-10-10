package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program sorts three numbers.")
	var a, b, c float64
	fmt.Println("Enter three numbers:")
	fmt.Scan(&a, &b, &c)
	if b > a && b > c && c > a {
		a, b, c = a, c, b
		fmt.Println("Numbers in a sorted order:", a, b, c)
	} else if b > a && b > c && a > c {
		a, b, c = c, a, b
		fmt.Println("Numbers in a sorted order:", a, b, c)
	} else if c > a && c > b && a > b {
		a, b, c = b, a, c
		fmt.Println("Numbers in a sorted order:", a, b, c)
	}
}
