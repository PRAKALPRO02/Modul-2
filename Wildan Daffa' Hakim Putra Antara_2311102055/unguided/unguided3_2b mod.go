package main

import (
	"fmt"
	"math"
)

func main() {
	var berat1, berat2 float64

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong : ")
		fmt.Scan(&berat1, &berat2)
		if !((berat1+berat2 > 150) || berat1 < 0 || berat2 < 0) {
			fmt.Println("Sepeda motor pak andi akan oleng :", math.Abs(berat1-berat2) >= 9)
		} else {
			fmt.Println("Proses Selesai")
			break
		}
	}
}
