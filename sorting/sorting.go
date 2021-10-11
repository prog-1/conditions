package main

import "fmt"

func main() {
	fmt.Println("The program sorts three numbers.")
	var a, b, c float64
	fmt.Println("Enter three numbers:")
	fmt.Scan(&a, &b, &c)
	if a > b {
		a, b = b, a
		fmt.Println("Numbers in a sorted order:", a, b, c)
	} else if b > c {
		b, c = c, b
		fmt.Println("Numbers in a sorted order:", a, b, c)
	} else if a > c {
		a, c = c, a
		fmt.Println("Numbers in a sorted order:", a, b, c)
	}
}
