package main

import "fmt"

func main() {
	vsr()
}

func vsr() {
	var v1, v2 int
	var c int

	fmt.Scanln(&c)

	for i := 0; i < c; i++ {
		fmt.Scanln(&v1, &v2)
		fmt.Println((2 * v1 * v2) / (v1 + v2))
	}
}
