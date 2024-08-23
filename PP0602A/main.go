package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	pp0602a()
}

func pp0602a() {
	reader := bufio.NewReader(os.Stdin)
	var str string
	var c int
	fmt.Scanln(&c)

	for i := 0; i < c; i++ {
		str, _ = reader.ReadString('\n')
		strs := strings.Fields(str)

		for j := 1; j < len(strs); j++ {
			if j%2 == 0 {
				fmt.Print(strs[j], " ")
			}
		}

		for j := 1; j < len(strs); j++ {
			if j%2 != 0 {
				fmt.Print(strs[j], " ")
			}
		}

		fmt.Println()
	}
}
