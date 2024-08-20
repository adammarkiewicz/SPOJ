package main

import (
	"fmt"
	"math"
)

func main() {
	rownanie()
}

func rownanie() {
	var a, b, c, d float64

	for {
		_, err := fmt.Scanln(&a, &b, &c)
		if err != nil {
			break
		}

		d = math.Pow(b, 2) - 4*(a*c)

		if d == 0 {
			fmt.Println(1)
		} else if d < 0 {
			fmt.Println(0)
		} else {
			fmt.Println(2)
		}
	}
}
