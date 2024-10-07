package main

import (
	"fmt"
	"math"
)

func main() {
	var kantonga, kantongb float64
	var oleng bool
	for {
		fmt.Print("Masukkan berat belanjaan kedua kantong: ")
		fmt.Scan(&kantonga, &kantongb)
		oleng = math.Abs(kantonga-kantongb) >= 9

		if kantonga+kantongb > 150 {
			fmt.Print("Proses selesai")
			break
		}
		fmt.Println("Sepeda motor pak Andi akan oleng : ", oleng)
	}
}
