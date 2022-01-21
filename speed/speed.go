package main

import "fmt"

func main() {

	fmt.Println(`This program performs km/h <-> m/s speed conversion. Please enter
	(1) if you want to convert km/h to m/s;
	(2) if you want to convert m/s to km/h.`)

	var mode float64

	fmt.Scan(&mode)

	if mode == 1 {
		fmt.Print("Enter the speed in km/h: ")
		var speedinkmh float64
		fmt.Scan(&speedinkmh)
		speed := speedinkmh / 3.6
		fmt.Println("The speed in m/s is:", speed)
	} else if mode == 2 {
		fmt.Print("Enter the speed in m/s: ")
		var speedinms float64
		fmt.Scan(&speedinms)
		secondspeed := speedinms * 3.6
		fmt.Println("The speed in km/h is:", secondspeed)
	} else {
		fmt.Println("Please select one of the modes above!")
	}
}
