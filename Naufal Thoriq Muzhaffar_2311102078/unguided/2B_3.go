package main

import (
	"fmt"
	"math"
)

func main() {
	var kantong1, kantong2 float64
	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)
		if kantong1+kantong2 > 150 || kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}
		diff := math.Abs(kantong1 - kantong2)
		if diff >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
