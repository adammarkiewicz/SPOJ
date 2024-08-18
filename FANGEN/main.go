package main

import (
	"fmt"
)

func main() {
	fangen()
}

func fangen() {
	for {
		var fanRadius, fanDiameter int
		fmt.Scanln(&fanRadius)

		if fanRadius == 0 {
			break
		}

		if fanRadius > 0 {
			fanDiameter = 2 * fanRadius
		} else {
			fanDiameter = -2 * fanRadius
		}

		var fanArray = make([][]rune, fanDiameter)
		for i := 0; i < fanDiameter; i++ {
			fanArray[i] = make([]rune, fanDiameter)
		}

		for i := 0; i < fanDiameter/2; i++ {
			for j := 0; j < fanDiameter/2; j++ {
				if fanRadius > 0 {
					if j > i {
						fanArray[i][j] = '.'
					} else {
						fanArray[i][j] = '*'
					}
				} else {
					if j < i {
						fanArray[i][j] = '.'
					} else {
						fanArray[i][j] = '*'
					}
				}
			}
		}

		for i := 0; i < fanDiameter/2; i++ {
			for j := fanDiameter / 2; j < fanDiameter; j++ {
				fanArray[i][j] = fanArray[(fanDiameter/2)-(i+1)][j-(fanDiameter/2)]
			}
		}

		for i := fanDiameter / 2; i < fanDiameter; i++ {
			for j := fanDiameter / 2; j < fanDiameter; j++ {
				fanArray[i][j] = fanArray[j-(fanDiameter/2)][i-(fanDiameter/2)]
			}
		}

		for i := fanDiameter / 2; i < fanDiameter; i++ {
			for j := 0; j < fanDiameter/2; j++ {
				fanArray[i][j] = fanArray[i-(fanDiameter/2)][(fanDiameter/2)-(j+1)]
			}
		}

		for i := 0; i < fanDiameter; i++ {
			for j := 0; j < fanDiameter; j++ {
				fmt.Print(string(fanArray[i][j]))
			}
			fmt.Println()
		}

		fmt.Println()
	}
}
