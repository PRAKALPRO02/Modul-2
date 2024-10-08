package main

import (
	"fmt"
	"math"
)

func main() {
	var kantong1, kantong2 float64

	for {
		// Meminta input berat belanjaan
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		// Jika berat salah satu kantong negatif, keluar dari loop
		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Jika total berat kedua kantong lebih dari 150 kg, keluar dari loop
		if kantong1+kantong2 > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Hitung selisih berat kedua kantong
		selisih := math.Abs(kantong1 - kantong2)

		// Tampilkan apakah sepeda motor akan oleng
		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
