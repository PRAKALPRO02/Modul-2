package main

import (
	"fmt"
	"math"
)

func main() {
	var berat1, berat2 float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		// Hentikan program jika salah satu kantong negatif
		if berat1 < 0 || berat2 < 0 {
			break
		}

		// Hentikan program jika total berat melebihi 150 kg
		if berat1+berat2 > 150 {
			break
		}

		// Periksa selisih berat kedua kantong
		if math.Abs(berat1-berat2) >= 9 {
			fmt.Println("true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng : false")
		}
	}

	fmt.Println("Proses selesai.")
}
