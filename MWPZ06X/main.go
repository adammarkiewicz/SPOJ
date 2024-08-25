package main

import (
	"fmt"
	"math"
)

func main() {
	mwpz06x()
}

func mwpz06x() {
	var n float64
	var c int
	fmt.Scanln(&c)

	for i := 0; i < c; i++ {
		fmt.Scanln(&n)

		fmt.Println(math.Pow(n, 2))
	}
}
