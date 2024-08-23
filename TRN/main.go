package main

import "fmt"

func main() {
	trn()
}

func trn() {
	var i, j int
	fmt.Scanln(&i, &j)

	tranMx := make([][]int, j)
	for y := 0; y < j; y++ {
		tranMx[y] = make([]int, i)
	}

	for x := 0; x < i; x++ {
		for y := 0; y < j; y++ {
			fmt.Scan(&tranMx[y][x])
		}
	}

	for y := 0; y < j; y++ {
		for x := 0; x < i; x++ {
			fmt.Print(tranMx[y][x], " ")
		}

		fmt.Println()
	}
}
