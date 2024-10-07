package main

import (
	"fmt"
)

// Fungsi untuk menghitung biaya pengiriman berdasarkan berat
func hitungBiayaPengiriman(beratGram int) (int, int, int) {
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	biayaPerKg := 10000
	biayaTambahan := calculateAdditionalCost(sisaGram, beratKg)

	totalBiaya := (beratKg * biayaPerKg) + biayaTambahan
	return beratKg, sisaGram, totalBiaya
}

// Fungsi untuk menghitung biaya tambahan berdasarkan sisa berat
func calculateAdditionalCost(sisaGram int, beratKg int) int {
	if beratKg > 10 {
		return 0
	} else if sisaGram >= 500 {
		return sisaGram * 5
	}
	return sisaGram * 15
}

func main() {
	var beratParcel int
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&beratParcel)

	beratKg, sisaGram, biaya := hitungBiayaPengiriman(beratParcel)
	fmt.Printf("Berat parsel (gram): %d\n", beratParcel)
	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", beratKg*10000, biaya-(beratKg*10000))
	fmt.Printf("Total biaya: Rp. %d\n", biaya)
}
