package main

import "fmt"

func main() {
	skarbfi()
}

func skarbfi() {
	var c int
	fmt.Scanln(&c)

	for i := 0; i < c; i++ {
		var north, south, west, east int
		var directionNorthSouth, stepsNortSouth, directionWestEast, stepsWestEast int

		var c2 int
		fmt.Scanln(&c2)

		for j := 0; j < c2; j++ {

			var d, s int
			fmt.Scanln(&d, &s)

			if d == 0 {
				north += s
			} else if d == 1 {
				south += s
			} else if d == 2 {
				west += s
			} else if d == 3 {
				east += s
			}
		}

		northSouth := north - south
		westEast := west - east

		if northSouth == 0 && westEast == 0 {
			fmt.Println("studnia")
			continue
		}

		if northSouth > 0 {
			directionNorthSouth = 0
			stepsNortSouth = northSouth
		} else {
			directionNorthSouth = 1
			stepsNortSouth = -northSouth
		}

		if westEast > 0 {
			directionWestEast = 2
			stepsWestEast = westEast
		} else {
			directionWestEast = 3
			stepsWestEast = -westEast
		}

		if stepsNortSouth == 0 {
			fmt.Println(directionWestEast, stepsWestEast)
		} else if stepsWestEast == 0 {
			fmt.Println(directionNorthSouth, stepsNortSouth)
		} else {
			fmt.Println(directionNorthSouth, stepsNortSouth)
			fmt.Println(directionWestEast, stepsWestEast)
		}
	}
}
