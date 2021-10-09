package main

import "fmt"

func main() {
	fmt.Println("The program determines if a year is a leap year.")
	fmt.Println("Enter the year:")
	var year uint
	fmt.Scanln(&year)
	if year%100 == 0 {
		if year%400 == 0 {
			fmt.Println(year, "is a leap year.")
		} else {
			fmt.Println(year, "is not a leap year.")
		}
	} else if year%100 != 0 {
		if year%4 == 0 {
			fmt.Println(year, "is a leap year.")
		} else {
			fmt.Println(year, "is not a leap year.")
		}
	}
}
