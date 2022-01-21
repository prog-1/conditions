package main

import "fmt"

func main() {
	fmt.Println("The program determines if a year is a leap year.")
	fmt.Println("Enter the Year:")
	var y int
	fmt.Scan(&y)
	if y%100 == 0 && y%40 == 0 {
		fmt.Println(y, "leap year")
	} else {
		fmt.Println(y, "not leap")
		fmt.Println(y, "not leap")
	}
}
