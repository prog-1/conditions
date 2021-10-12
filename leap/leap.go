package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a year is a leap year.")
	fmt.Println("Enter the year:")
	var year int 
	fmt.Scanln(&year)
	if year%100 == 0 {
		if year%400 == 0 {
			fmt.Println("Is a leap year", year)
		} else {
			fmt.Println("Is not a leap year", year)
		}
	} else {
		if year%4 == 0 {
			fmt.Println("Is a leap year", year)
		} else {
			fmt.Println("Is not a leap year", year)
		}
	}
}