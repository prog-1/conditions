package main

import "fmt"

func main() {
	fmt.Print("The program determines if a year is a leap year.")
	var a uint
	fmt.Println("Enter year")

	fmt.Scanln(&a)
	if a%400 == 0 {
		fmt.Println(a, "is a leap year")
		return
	} else {
		if a%100 == 0 {
			fmt.Println(a, "is not a leap year")
		} else {
			if a%4 == 0 {
				fmt.Println(a, "is a leap year")

			} else {
				fmt.Println(a, "is not a leap year")
			}
		}
	}
}
