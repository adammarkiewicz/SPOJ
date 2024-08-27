package main

import "fmt"

func main() {
	var pesel string
	var c, checkSum int
	fmt.Scanln(&c)

	for i := 0; i < c; i++ {
		fmt.Scanln(&pesel)

		checkSum = int(pesel[0]-'0') + (int(pesel[1]-'0') * 3) + (int(pesel[2]-'0') * 7) +
			(int(pesel[3]-'0') * 9) + int(pesel[4]-'0') + (int(pesel[5]-'0') * 3) +
			(int(pesel[6]-'0') * 7) + (int(pesel[7]-'0') * 9) + int(pesel[8]-'0') +
			(int(pesel[9]-'0') * 3) + int(pesel[10]-'0')

		if checkSum%10 == 0 {
			fmt.Println("D")
		} else {
			fmt.Println("N")
		}
	}
}
