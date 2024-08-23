package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sys()
}

func sys() {
	var c int
	var l int64

	fmt.Scanln(&c)

	for i := 0; i < c; i++ {
		fmt.Scanln(&l)

		fmt.Println(strings.ToUpper(strconv.FormatInt(l, 16)), strings.ToUpper(strconv.FormatInt(l, 11)))
	}
}

// func toHex(n int) string {
// 	var hex string
// 	hexArr := make([]rune, 100)
// 	i := 0
// 	for n > 0 {
// 		rem := n % 16

// 		if rem < 10 {
// 			hexArr[i] = rune(rem + 48)
// 		} else {
// 			hexArr[i] = rune(rem + 55)
// 		}

// 		n /= 16
// 		i++
// 	}

// 	for j := i - 1; j >= 0; j-- {
// 		hex += string(hexArr[j])
// 	}

// 	return hex
// }
