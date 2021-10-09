package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Println("Enter the hole dimensions (2 numbers):")
	fmt.Scan(&a, &b)
	var p, q, r float64
	fmt.Println("Enter the brick dimensions (3 numbers):")
	fmt.Scan(&p, &q, &r)
	if p < r && p > q && r > q {
		p, r, q = q, p, r
	} else if r < p && q < p && r < q {
		p, r, q = r, q, p
	} else if p < q && q < r && p < r {
		r, q = q, r
	} else if q < r && r < p && q < p {
		p, q = q, p
	} else if r < p && p < q && r < q {
		p, r = r, p
	}
	if p <= a && r <= b || p <= b && r <= a {
		if p <= a && q <= b || p < b && q <= a {
			if q <= a && r <= b || q <= b && r <= a {
				fmt.Println("The brick can be inserted in the hole!")
			} else {
				fmt.Println("The brick can be inserted in the hole!")
			}
		} else {
			fmt.Println("The brick can be inserted in the hole!")
		}
	} else {
		fmt.Println("The brick cannot be inserted in the hole.")
	}
}
