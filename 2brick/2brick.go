package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d, e float64
	fmt.Println("Enter the hole dimensions (2 numbers):")
	fmt.Scan(&a, &b)
	fmt.Println("Enter the brick dimensions (3 numbers):")
	fmt.Scan(&c, &d, &e)
	max2 := (d + e + math.Abs(a-b)) / 2
	max := (c + max2 + math.Abs(c-max2)) / 2
	min2 := (d + e - math.Abs(a-b)) / 2
	min := (c + min2 - math.Abs(c-min2)) / 2
	if d != min && d != max {
		if (a >= min && b >= d) || (a >= d && b >= min) {
			fmt.Println("The brick can fit in the hole!") // At first, I've just made a variable, but there was a mistake "declared and not used" and I removed the variable.
		} else {
			fmt.Println("The brick can't fit in the hole.")
		}
	} else {
		if e != min && e != max {
			if (a >= min && b >= c) || (a >= c && b >= min) {
				fmt.Println("The brick can fit in the hole!")
			} else {
				fmt.Println("The brick can't fit in the hole.")
			}
		} else {
			if (a >= min && b >= d) || (a >= d && b >= min) {
				fmt.Println("The brick can fit in the hole!")
			} else {
				fmt.Println("The brick can't fit in the hole.")
			}

		}
	}
} // I'm not sure that this is what you wanted, but anyway it's something different
