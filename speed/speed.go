package main

import "fmt"

func main() {
	fmt.Println("This program performs km/h <-> m/s speed conversion. Please enter")
	fmt.Println("(1) if you want to convert km/h to m/s;")
	fmt.Println("(2) if you want to convert m/s to km/h;")
	var a float64
	fmt.Scan(&a)
	if a == 1 {
		fmt.Println("Enter km/h:")
		var b float64
		fmt.Scan(&b)
		fmt.Println("result:", b*5/18)
	} else if a == 2 {
		fmt.Println("Enter m/s:")
		var b float64
		fmt.Scan(&b)
		fmt.Println("result:", b*18/5)
	} else {
		fmt.Println("Error")
	}

}
