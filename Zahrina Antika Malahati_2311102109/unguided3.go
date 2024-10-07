package main

import (
	"fmt"
	"math"
)

func main() {
	var berat1, berat2 float64
	var totalBerat float64
	var selisihBerat float64

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scanf("%f %f", &berat1, &berat2)

		// Menghitung selisih berat antara kedua kantong
		selisihBerat = math.Abs(berat1 - berat2)
		totalBerat = berat1 + berat2

		// Mengecek apakah selisih berat lebih dari 9 kg
		if selisihBerat >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}

		// Mengecek apakah berat total melebihi 150 kg atau salah satu kantong beratnya negatif
		if totalBerat > 150 || berat1 < 0 || berat2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}

//Zahrina Antika Malahati_2311102109
