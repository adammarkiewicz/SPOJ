package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()

	var inpTxt string
	x, y, z := new(big.Int), new(big.Int), new(big.Int)

	for scanner.Scan() {
		inpTxt = scanner.Text()
		x.SetString(inpTxt, 10)

		scanner.Scan()
		inpTxt = scanner.Text()
		y.SetString(inpTxt, 10)

		z.Mul(x, y)
		fmt.Println(z.String())
	}
}
