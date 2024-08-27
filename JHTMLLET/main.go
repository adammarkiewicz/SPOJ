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
			if runes[i] == '<' {
				for runes[i] != '>' {
					i++
					runes[i] = unicode.ToUpper(runes[i])
				}
			}
		}

		fmt.Println(string(runes))
	}
}
