package main

import "fmt"

func main() {
	pp0601b()
}

func pp0601b() {
	var c, n, x, y int

	fmt.Scanln(&c)

	for i := 0; i < c; i++ {
		fmt.Scanln(&n, &x, &y)

		for j := 1; j < n; j++ {
			if j%x == 0 && j%y != 0 {
				fmt.Print(j, " ")
			}
		}

		fmt.Println()
	}
}
