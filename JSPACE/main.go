package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		for i := 0; i < len(runes); i++ {
			if runes[i] == ' ' {
				runes[i] = '0'
				runes[i+1] = unicode.ToUpper(runes[i+1])
			}
		}

		for i := 0; i < len(runes); i++ {
			if runes[i] != '0' {
				fmt.Print(string(runes[i]))
			}
		}

		fmt.Println()
	}
}
