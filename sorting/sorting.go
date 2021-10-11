package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program sorts three numbers.")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	if a < b && a < c {
		if b < c {
			fmt.Println("Number in a sorted order:", a, b, c)
		} else {
			fmt.Println("Number in a sorted order:", a, c, b)
		}
	} else if b < a && b < c {
		if a < c {
			fmt.Println("Number in a sorted number:", b, a, c)
		} else {
			fmt.Println("Numbers in a sorted order:", b, c, a)
		} 
	} else if c < a && c < b {
		if a < b {
			fmt.Println("Numbers in a sorted order:", c, a, b)
		} else {
			fmt.Println("Numbers in a sorted order:", c, b, a)
		}
	} else {
		fmt.Println("All numbers are equal.")
	}
}