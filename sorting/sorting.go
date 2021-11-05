package main

import "fmt"

func main() {
	fmt.Println("program sorts three numbers.")
	fmt.Println("Enter the numbers:")
	var e, f, g float64
	fmt.Scanln(&e, &f, &g)
	if e < f && e < g && f < g {
		fmt.Println("Numbers in a sorted order:", e, f, g)
	} else {
		fmt.Println("Numbers in a sorted order:", g, f, e)
	}
}
