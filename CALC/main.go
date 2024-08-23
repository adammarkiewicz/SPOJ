package main

import "fmt"

func main() {
	calc()
}

func calc() {
	var op string
	var i1, i2 int
	for {
		_, err := fmt.Scanln(&op, &i1, &i2)
		if err != nil {
			break
		}

		if op == "+" {
			fmt.Println(i1 + i2)
		} else if op == "-" {
			fmt.Println(i1 - i2)
		} else if op == "*" {
			fmt.Println(i1 * i2)
		} else if op == "/" {
			fmt.Println(i1 / i2)
		} else {
			fmt.Println(i1 % i2)
		}
	}
}
