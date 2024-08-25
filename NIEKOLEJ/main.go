package main

import (
	"fmt"
)

func main() {
	niekolej()
}

func niekolej() {
	var n int
	fmt.Scanln(&n)

	if n == 0 {
		fmt.Println(0)
		return
	}
	if n <= 2 {
		fmt.Println("NIE")
		return
	}

	for j := 1; j <= n; j += 2 {
		fmt.Print(j, " ")
	}

	for j := 0; j <= n; j += 2 {
		fmt.Print(j, " ")
	}

	fmt.Println()
}
