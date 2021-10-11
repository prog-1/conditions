package main

import "fmt"

func main() {
	fmt.Println("The program sorts three numbers.")
	fmt.Println("Enter the numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if a > b && b > c {
		fmt.Println(c, b, a)
	} else if c > b && b > a {
		fmt.Println(a, b, c)
	} else if b > a && a > c {
		fmt.Println(c, a, b)
	} else if c > a && a > b {
		fmt.Println(b, a, c)
	} else if a > c && c > b {
		fmt.Println(b, c, a)
	} else if b > c && c > a {
		fmt.Println(a, c, b)
	}
}
