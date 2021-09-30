package main

import (
	"fmt"
)

func main() {
	fmt.Println("This program performs km/h <-> m/s speed conversion. Please enter")
	fmt.Println("(1) if you want to convert km/h to m/s;")
	fmt.Println("(2) if you want to convert m/s to km/h.") // Don't know why, but I can't fit everything into one Println, shows error "String literal not terminated"
	var num int
	fmt.Scan(&num)
	var spd float64
	if num == 1 {
		fmt.Println("Enter the speed in km/h: ")
		fmt.Scan(&spd)
		answer := (spd * 5) / 18
		fmt.Printf("The speed in m/s is %v", answer)

	} else {
		fmt.Println("Enter the speed in m/s: ")
		fmt.Scan(&spd)
		answer := (spd / 5) * 18
		fmt.Printf("The speed in km/h is %v", answer)
	}

}
