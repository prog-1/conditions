package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a year is a leap year.")
	var year insertedfmt.Println("Enter the year:")
	fmt.Scan(&year)
	if year%4 == 0 {
		if year%400 == 0 {
			fmt.Println(year, "is a leap year.")
		} else {
			fmt.Println(year, "is a not leap year.")
		}
	} else if year%4 ==0 {
		fmt.Println(year, "is a leap year.")
	} else {
		fmt.Println(year, "is a not leap year.")
	}
}