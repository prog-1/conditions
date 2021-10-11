package main

import (
	"fmt"
)

func main() {
	fmt.Println("This program performs km/h <-> m/s speed conversion. Please enter
	(1) if you want to convert km/h to m/s;
	(2) if you want to convert m/s to km/h.")
	var mode float64
	fmt.Scanln(&mode)
	if mode == 1 {
		fmt.Println("Enter the speed in km/h:")
		var speedinkmh float64
		fmt.Scanln(&speedinkmh)
		speed := speedinkmh / 3.6
		fmt.Println("The speed in m/s:", speed)
	} else {
		if mode == 2 {
			fmt.Println("Enter the speed in m/s:")
			var speedinms float64
			fmt.Scanln(&speedinms)
			speed2 := speedinms * 3.6
			fmt.Println("The speed in km/h:", speed2)
		} else {
			fmt.Println("Wrong number")
		}
	}
}