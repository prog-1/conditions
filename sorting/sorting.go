package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program sorts three numbers.")
	fmt.Println("Enter the numbers:")
	var a, b, c float64
	var max, min, scnd float64
	fmt.Scan(&a, &b, &c)
	if a != b && a != c && b != c {
		if a > b && a > c {
			max = a
		}
		if b > a && b > c {
			max = b
		}
		if c > a && c > b {
			max = c
		}
		if a < b && a < c {
			min = a
		}
		if b < a && b < c {
			min = b
		}
		if c < a && c < b {
			min = c
		}
		scnd = a + b + c - (max + min)

		fmt.Printf("Numbers in a sorted order: %v %v %v", min, scnd, max)
	}
	if a == b && a == c && b == c {
		fmt.Printf("Numbers in a sorted order: %v %v %v", a, b, c)
	}
	if a == b || a == c || b == c {
		if a >= b && a >= c {
			max = a
		}
		if b >= a && b >= c {
			max = b
		}
		if c >= a && c >= b {
			max = c
		}
		if a <= b && a <= c {
			min = a
		}
		if b <= a && b <= c {
			min = b
		}
		if c <= a && c <= b {
			min = c
		}
		scnd = a + b + c - (max + min)

		fmt.Printf("Numbers in a sorted order: %v %v %v", min, scnd, max)
	}
}
