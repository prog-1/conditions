package main

import (
	"fmt"
)

func main() {
	fmt.Println(`This program performs km/h <-> m/s speed conversion. Please enter
	(1) if you want to convert km/h to m/s;
	(2) if you want to convert m/s to km/h.`)
	var sel float64
	fmt.Scan(&sel)
	if sel == 1 {
		fmt.Println("Enter the speed in km/h:")
		var kmh float64
		fmt.Scan(&kmh)
		ms := kmh / 3.6
		fmt.Println("The speed in m/s is", ms)
	} else if sel == 2 {
		fmt.Println("Enter the speed in m/s:")
		var mps float64
		fmt.Scan(&mps)
		kmph := mps * 3.6
		fmt.Println("The speed in km/h is", kmph)
	}
}
