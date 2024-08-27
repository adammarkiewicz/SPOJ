package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	charCount := make(map[rune]int)

	var sortedKeys []rune
	for ch := 'a'; ch <= 'z'; ch++ {
		sortedKeys = append(sortedKeys, ch)
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		sortedKeys = append(sortedKeys, ch)
	}

	var c int
	fmt.Scanln(&c)

	var line string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for i := 0; i < c; i++ {
		scanner.Scan()
		line = scanner.Text()

		for _, char := range line {
			charCount[char]++
		}
	}

	for _, key := range sortedKeys {
		if charCount[key] != 0 {
			fmt.Printf("%s %d\n", string(key), charCount[key])
		}
	}
}
