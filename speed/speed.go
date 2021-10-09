package main

import "fmt"

func main() {
	fmt.Println("This program performs km/h <-> m/s speed conversion. Please enter")
	fmt.Println("(1) if you want to convert km/h to m/s;")
	fmt.Println("(2) if you want to convert m/s to km/h.")
	var number float64
	fmt.Scanln(&number)
	if number == 1 {
		fmt.Println("Enter the speed in km/h:")
		var speedkmh float64
		fmt.Scanln(&speedkmh)
		speedms := speedkmh / 3.6
		fmt.Println("The speed in m/s is:", speedms)
	} else if number == 2 {
		fmt.Println("Enter the speed in m/s:")
		var speedms float64
		fmt.Scanln(&speedms)
		speedkmh := speedms * 3.6
		fmt.Println("The speed in km/h is", speedkmh)
	} else {
		fmt.Println("Error")
	}
}
