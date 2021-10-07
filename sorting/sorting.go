package main

import "fmt"

func main() {
	fmt.Println("The program sorts three numbers.")
	fmt.Print("Enter the numbers: ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if a < b && a < c {
		if b < c {
			fmt.Println("Numbers in a sorted order:", a, b, c)
		} else {
			fmt.Println("Numbers in a sorted order:", a, c, b)
		}
	} else if b < a && b < c {
		if a < c {
			fmt.Println("Numbers in a sorted order:", b, a, c)
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
		fmt.Println("ALL numbers are equal.")
	}
}
