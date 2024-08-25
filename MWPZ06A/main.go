package main

import (
	"fmt"
	"math"
)

func main() {
	mwpz06a()
}

func mwpz06a() {
	var x, y, z float64
	var c int
	fmt.Scanln(&c)

	for i := 0; i < c; i++ {
		fmt.Scanln(&x, &y, &z)

		fmt.Println(math.Round((x + y - (y * z)) / (1 - z) * 12))
	}
}
