package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program determines if a year is a leap year.")
	fmt.Println("Enter the year:")
	var a int64
	fmt.Scan(&a)
	if a%4 == 0 {
		if a%100 == 0 {
			if a%400 == 0 {
				fmt.Println(a, "is a leap year.")
			} else {
				fmt.Println(a, "is a not leap year.")
			}
		} else {
			fmt.Println(a, "is a leap year.")
		}
	} else {
		fmt.Println(a, "is a not leap year.")
	}

}
