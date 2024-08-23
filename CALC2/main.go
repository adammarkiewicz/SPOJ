package main

import "fmt"

func main() {
	calc2()
}

func calc2() {
	var r []int = make([]int, 10)
	var op string
	var i1, i2 int
	for {
		_, err := fmt.Scanln(&op, &i1, &i2)
		if err != nil {
			break
		}

		if op == "z" {
			r[i1] = i2
		} else if op == "+" {
			fmt.Println(r[i1] + r[i2])
		} else if op == "-" {
			fmt.Println(r[i1] - r[i2])
		} else if op == "*" {
			fmt.Println(r[i1] * r[i2])
		} else if op == "/" {
			fmt.Println(r[i1] / r[i2])
		} else {
			fmt.Println(r[i1] % r[i2])
		}
	}
}
