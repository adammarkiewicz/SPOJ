package main

import (
	"fmt"
	"math"
)

func main() {
	glutton()
}

func glutton() {
	var c int
	fmt.Scanln(&c)

	for i := 0; i < c; i++ {
		var gluttonCount, cookiesInBox, neededCookies int
		fmt.Scanln(&gluttonCount, &cookiesInBox)

		for j := 0; j < gluttonCount; j++ {
			var eatingSpeed int
			fmt.Scanln(&eatingSpeed)

			neededCookies += 86400 / eatingSpeed
		}

		fmt.Println(math.Ceil(float64(neededCookies) / float64(cookiesInBox)))
	}
}
