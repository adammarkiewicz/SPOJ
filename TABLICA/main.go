package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tablica()
}

func tablica() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	s := strings.Fields(text)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Print(strings.Join(s, " "))
}
