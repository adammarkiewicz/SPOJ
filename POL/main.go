package main

import "fmt"

func main() {
	pol()
}

func pol() {
	var c int
	var s string

	fmt.Scanln(&c)

	for i := 0; i < c; i++ {
		fmt.Scanln(&s)
		fmt.Println(s[0 : len(s)/2])
	}
}
