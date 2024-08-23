package main

import "fmt"

func main() {
	trn()
}

func trn() {
	var i, j int
	fmt.Scanln(&i, &j)

	mx := make([]int, i*j)

	for x := 0; x < i; x++ {
		for y := 0; y < j; y++ {
			fmt.Scan(&mx[j*x+y])
		}
	}

	for y := 0; y < j; y++ {
		for x := 0; x < i; x++ {
			fmt.Print(mx[j*x+y], " ")
		}

		fmt.Println()
	}
}
