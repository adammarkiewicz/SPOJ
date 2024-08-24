package main

import "fmt"

func main() {
	pp0602b()
}

func pp0602b() {
	var arr [101][101]int
	var tmp [2]int
	var c, l, k int
	fmt.Scanln(&c)

	for z := 0; z < c; z++ {
		fmt.Scanln(&l, &k)

		for x := 0; x < l; x++ {
			for y := 0; y < k; y++ {
				fmt.Scan(&arr[x][y])
			}
		}

		tmp[0] = arr[0][0]
		tmp[1] = arr[l-1][k-1]

		for i := 0; i < k-1; i++ {
			arr[0][i] = arr[0][i+1]
		}
		for i := k - 1; i > 0; i-- {
			arr[l-1][i] = arr[l-1][i-1]
		}
		for i := l - 1; i > 0; i-- {
			arr[i][0] = arr[i-1][0]
		}
		for i := 0; i < l-1; i++ {
			arr[i][k-1] = arr[i+1][k-1]
		}

		arr[1][0] = tmp[0]
		arr[l-2][k-1] = tmp[1]

		for a := 0; a < l; a++ {
			for b := 0; b < k; b++ {
				fmt.Print(arr[a][b], " ")
			}

			fmt.Print("\n")
		}
	}
}
