package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a year is a leap year.")
	var y int
	fmt.Println("Enter the year:")
	fmt.Scan(&y)
	if y%100 == 0 {
	} else if y%400 == 0 {
		fmt.Println(y, "is a leap year")
		if y%100 != 0 {
		} else if y%400 != 0 {
			fmt.Println(y, "is not a leap year")
		}
	} else if y%4 == 0 {
		fmt.Println(y, "is a leap year")
	} else if y%4 != 0 {
		fmt.Println(y, "is a not a leap year")
	}
}
