package main

import (
	"fmt"
	"math/big"
)

func main() {
	binoms()
}

func binoms() {
	var c int
	fmt.Scanln(&c)

	result, n, k := big.NewInt(0), big.NewInt(0), big.NewInt(0)
	for i := 0; i < c; i++ {
		fmt.Scanln(n, k)

		result.Div(fac(n), result.Mul(fac(result.Sub(n, k)), fac(k)))
		fmt.Println(result.String())
	}
}

var facMemo = [10000]*big.Int{}

func fac(x *big.Int) *big.Int {
	i := x.Uint64()
	if facMemo[i] != nil {
		return facMemo[i]
	} else {
		n := big.NewInt(1)
		if x.Cmp(big.NewInt(0)) == 0 {
			return n
		}
		facMemo[i] = n.Mul(x, fac(n.Sub(x, n)))
		return facMemo[i]
	}
}
