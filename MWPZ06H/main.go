package main

import (
	"fmt"
	"sort"
)

func main() {
	mwpz06h()
}

func mwpz06h() {
	var t, l int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		fmt.Scanln(&l)
		sl := make([]int, l)

		for j := 0; j < l; j++ {
			fmt.Scan(&sl[j])
		}

		sort.Sort(sort.Reverse(sort.IntSlice(sl)))
		maxCnt := maxCount(sl)
		for k := 0; k < maxCnt; k++ {
			fmt.Print(sl[0], " ")
		}

		for m := l - 1; m >= maxCnt; m-- {
			fmt.Print(sl[m], " ")
		}

		fmt.Println()
	}
}

func maxCount(slice []int) int {
	c := 1
	for i := 0; i < len(slice); i++ {
		if i+1 < len(slice) && slice[i] == slice[i+1] {
			c++
		} else {
			break
		}
	}

	return c
}
