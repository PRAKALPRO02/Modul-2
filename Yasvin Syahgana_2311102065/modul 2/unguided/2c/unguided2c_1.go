package main

import (
	"fmt"
)

func hitungBiayaKirim(beratGram int) int {
	// Konversi berat ke kg dan sisa gram
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	// Biaya dasar per kg
	biayaPerKg := 10000
	totalBiaya := beratKg * biayaPerKg

	// Ketentuan tambahan biaya untuk sisa berat
	var sisaBiaya int
	if beratKg > 10 {
		// Jika total berat lebih dari 10kg, sisa berat digratiskan
		sisaBiaya = 0
	} else {
		if sisaGram >= 500 {
			// Jika sisa lebih dari atau sama dengan 500 gram
			sisaBiaya = sisaGram * 5
		} else {
			// Jika sisa kurang dari 500 gram
			sisaBiaya = sisaGram * 15
		}
	}

	totalBiaya += sisaBiaya
	return totalBiaya
}

func main() {
	var beratParsial int

	fmt.Print("Masukkan berat parsel dalam gram: ")
	_, err := fmt.Scan(&beratParsial)
	if err != nil || beratParsial < 0 {
		fmt.Println("Input tidak valid. Harap masukkan angka yang benar dan tidak negatif.")
		return
	}

	biaya := hitungBiayaKirim(beratParsial)
	fmt.Printf("Biaya pengiriman untuk parsel seberat %d gram adalah: Rp. %d\n", beratParsial, biaya)
}