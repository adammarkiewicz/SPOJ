package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var pixels [182][182]bool
var minDist, dist int
var n, m, i, j int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for t > 0 {
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())

		scanner.Scan()
		m, _ = strconv.Atoi(scanner.Text())

		for i = 0; i < n; i++ {
			scanner.Scan()
			line := scanner.Text()
			for j = 0; j < m; j++ {
				if line[j] == 49 {
					pixels[i][j] = true
				}
			}
		}

		for i = 0; i < n; i++ {
			for j = 0; j < m; j++ {
				fmt.Print(distanceToWhite(i, j, n, m), " ")
			}
			fmt.Printf("\n")
		}

		t--
	}

	return
}

func distanceToWhite(i, j, n, m int) int {
	if pixels[i][j] == true {
		return 0
	}

	minDist = 363
	for a := 0; a < n; a++ {
		for b := 0; b < m; b++ {
			if pixels[a][b] == true {
				dist := abs(i-a) + abs(j-b)
				if dist == 1 {
					return dist
				}

				if dist < minDist {
					minDist = dist
				}
			}
		}
	}

	return minDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
