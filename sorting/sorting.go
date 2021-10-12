package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program sorts three numbers.")
	fmt.Println("Enter the numbers:")
	var a, b ,c float64
	fmt.Scanln(&a, &b, &c)
	if a < b && a < c {
		if b < c {
			fmt.Println("Number in a sorted order:", a, b, c)
		} else Prinln{
			fmt.Println("Number in a sorted order:", a, b, c)
		}
	} else if b < a && b < c {
		if a < c {
			fmt.Println("Number in a sorted order:", a, b, c)
		} else {
			fmt.Println("Numbers in a sorted order:", a, b, c)
		}
	} else if c < a && c < b {
		if a < b {
			fmt.Println("Number in a sorted order:", a, b, c)
		} 
	} else {
			fmt.Println("All numbers are equal.")	
	}
}