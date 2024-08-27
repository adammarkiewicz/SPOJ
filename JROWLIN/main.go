package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, x float64
	fmt.Scanln(&a, &b, &c)

	if a == 0 {
		if c-b == 0 {
			fmt.Println("NWR")
			return
		}

		fmt.Println("BR")
		return
	}

	x = (c - b) / a
	x = math.Round(x*100) / 100

	fmt.Printf("%.2f\n", x)
}
