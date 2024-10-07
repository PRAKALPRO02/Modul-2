package main

import (
	"fmt"
	"math"
)

func main() {
	var kantongKiri, kantongKanan float64

	for {
		fmt.Printf("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantongKiri, &kantongKanan)

		if !((kantongKiri+kantongKanan > 150) || kantongKiri < 0 || kantongKanan < 0) {
			fmt.Println("Sepeda motor pak Andi akan oleng :", math.Abs(kantongKiri-kantongKanan) >= 9)
		} else {
			fmt.Println("Proses selesai")
			break
		}
	}
}
