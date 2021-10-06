package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program sorts three numbers.")
	fmt.Println("Enter the numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if a < b && a < c {
		if b < c {
			fmt.Println("Numbers in a sorted order:", a, b, c)
		} else {
			fmt.Println("Numbers in a sorted order:", a, c, b)
		}
	} else {
		if a < b {
			fmt.Println("Numbers in a sorted order:", c, a, b)
		} else {
			if a < c {
				fmt.Println("Numbers in a sorted order:", b, a, c)
			} else {
				if b < c {
					fmt.Println("Numbers in a sorted order:", b, c, a)
				} else {
					fmt.Println("Numbers in a sorted order:", c, b, a)
				}
			}
		}
	}

}
