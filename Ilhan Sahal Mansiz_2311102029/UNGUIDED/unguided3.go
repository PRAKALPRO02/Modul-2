package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var beratKiri, beratKanan float64

		fmt.Print("Masukkan berat belanjaan di kedua kantong (kiri kanan): ")
		fmt.Scan(&beratKiri, &beratKanan)

		if beratKiri < 0 || beratKanan < 0 {
			fmt.Println("Salah satu kantong berisi berat negatif. Program dihentikan.")
			break
		}

		if beratKiri+beratKanan > 150 {
			fmt.Println("Total berat kedua kantong melebihi 150 kg. Program dihentikan.")
			break
		}

		selisih := math.Abs(beratKiri - beratKanan)

		if selisih >= 9 {
			fmt.Println("True, selisih berat kedua kantong lebih dari atau sama dengan 9 kg.")
		} else {
			fmt.Println("False, selisih berat kedua kantong kurang dari 9 kg.")
		}

		fmt.Println("Proses selesai.")
	}
}