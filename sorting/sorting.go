package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program sorts three numbers.")
	fmt.Print("Enter the numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if a > b {
		if b > c {
			fmt.Print("Numbers in a sorted order:", c, b, a)
			return
		} else {
			if b < c {
				if a < c {
					fmt.Print("Numbers in a sorted order:", b, a, c)
					return
				} else {
					fmt.Print("Numbers in a sorted order:", b, c, a)
					return
				}

			}
		}

	} else {
		if a < c {
			if b < c {
				fmt.Print("Numbers in a sorted order:", a, b, c)
				return
			} else {
				fmt.Print("Numbers in a sorted order:", a, c, b)
				return
			}
		} else {
			fmt.Print("Numbers in a sorted order:", c, a, b)
		}
	}

}
