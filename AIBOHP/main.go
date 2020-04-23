package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	memo [6101][6101]int
	str  string
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()

	for scanner.Scan() {
		memo = [6101][6101]int{}
		str = scanner.Text()
		fmt.Println(countInserts(0, len(str)-1))
	}

	return
}

func countInserts(x, y int) int {
	if x >= y {
		return 0
	}

	if x == y-1 {
		if str[x] == str[y] {
			return 0
		}
		return 1
	}

	if memo[x][y] != 0 {
		return memo[x][y]
	}

	if str[x] == str[y] {
		memo[x][y] = countInserts(x+1, y-1)
		return memo[x][y]
	}

	memo[x][y] = 1 + min(countInserts(x+1, y), countInserts(x, y-1))
	return memo[x][y]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
