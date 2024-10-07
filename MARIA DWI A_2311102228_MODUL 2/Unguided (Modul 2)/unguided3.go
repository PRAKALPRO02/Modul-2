package main

import (
	"fmt"
	"math"
)

func main() {
	var beratKiri, beratKanan float64

	for {
		
		fmt.Print("\nMasukkan berat belanjaan di kedua kantong: ")
		_, err := fmt.Scanln(&beratKiri, &beratKanan)

		
		if err != nil {
			fmt.Println("\nInput tidak valid. Mohon masukkan dua angka.")
			continue
		}

		
		if beratKiri < 0 || beratKanan < 0 {
			fmt.Println("\nProses selesai. Berat tidak bisa negatif.")
			break
		}

		
		totalBerat := beratKiri + beratKanan
		if totalBerat > 150 {
			fmt.Println("\nProses selesai. Total berat melebihi batas.")
			break
		}


		selisihBerat := math.Abs(beratKiri - beratKanan)
		akanOleng := selisihBerat >= 9

		fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", akanOleng)


		if beratKiri >= 90 || beratKanan >= 90 {
			fmt.Println("\n\nProses selesai. Berat salah satu kantong melebihi batas 90.\n")
			break
		}
	}

	
}
