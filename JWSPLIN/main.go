package main

import "fmt"

func main() {
	var c int
	fmt.Scanln(&c)

	var x1, y1, x2, y2, x3, y3 int
	for i := 0; i < c; i++ {
		fmt.Scanln(&x1, &y1, &x2, &y2, &x3, &y3)

		if (y2-y1)*(x3-x2) == (y3-y2)*(x2-x1) {
			fmt.Println("TAK")
		} else {
			fmt.Println("NIE")
		}
	}
}
