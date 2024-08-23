package main

import "fmt"

func main() {
	pp0601a2()
}

func pp0601a2() {
	var n, c, prev int
	c = 0
	prev = -1
	for {
		_, err := fmt.Scanln(&n)
		if err != nil {
			break
		}

		fmt.Println(n)

		if n == 42 && prev != 42 && prev != -1 {
			c++
			if c == 3 {
				break
			}
		}

		prev = n
	}
}
