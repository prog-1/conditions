package main

import "fmt"

func main() {
	fmt.Println("The program determines if a year is a leap year.")
	fmt.Print("Enter the year: ")
	var year int
	fmt.Scan(&year)
	if year%100 == 0 && year%4 == 0 {
		fmt.Println(year, "is a leap year.")
	} else {
		fmt.Println(year, "is not a leap year")
	}
	if year%400 == 0 {
		fmt.Println(year, "is a leap year.")
	} else {
		fmt.Println(year, "is a not leap year.")
	}
}
