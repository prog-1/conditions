package main

import "fmt"

func main() {
	fmt.Println("The program determines if a year is a leap year.")
	fmt.Print("Enter the year: ")
	var a int
	fmt.Scan(&a)
	if a%100 == 0 {
		if a%400 == 0 {
			fmt.Println(a, "is a leap year.")
		} else {
			fmt.Println(a, "is not a leap year.")
		}
	} else if a%4 == 0 {
		fmt.Println(a, "is a leap year.")
	} else {
		fmt.Println(a, "is not a leap year.")
	}

}
