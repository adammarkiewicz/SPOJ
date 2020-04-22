package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var pregen [1000000]int

func main() {
	for i := 0; i < 1000000; i++ {
		pregen[i] = exchangeCoin(i)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var coin int

	for scanner.Scan() {
		coin, _ = strconv.Atoi(scanner.Text())

		fmt.Println(exchangeCoin(coin))
	}

	return
}

func exchangeCoin(coin int) int {
	if coin < 1000000 && pregen[coin] != 0 {
		return pregen[coin]
	}

	if (coin/2)+(coin/3)+(coin/4) <= coin {
		return coin
	}
	return exchangeCoin(coin/2) + exchangeCoin(coin/3) + exchangeCoin(coin/4)
}
