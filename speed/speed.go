package main

import "fmt"

func main() {
	fmt.Println(`This program performs km/h <-> m/s speed conversion. Please enter
	(1) if you want to convert km/h to m/s;
	(2) if you want to convert m/s to km/h.`)
	fmt.Print(" ")
	var a float64
	fmt.Scan(&a)
	if a == 1 {
		fmt.Print("Enter the speed in m/s: ")
		fmt.Scan(&a)
		fmt.Println("The speed in km/h is ", (a * 3.6))
	} else if a == 2 {
		fmt.Print("Enter the speed in km/h: ")
		fmt.Scan(&a)
		fmt.Println("The speed in m/s is ", (a / 3.6))
	} else {
		fmt.Println("ERROR: Try typing 1 or 2;")
	}
	return
}
