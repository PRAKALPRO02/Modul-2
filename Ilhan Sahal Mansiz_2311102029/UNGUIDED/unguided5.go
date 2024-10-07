package main

import (
	"fmt"
)

func hitungBiayaPengiriman(beratGram int) (int, int, int) {
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	biayaPerKg := 10000
	biayaTambahan := 0

	if beratKg > 10 {
		biayaTambahan = 0
	} else if sisaGram >= 500 {
		biayaTambahan = sisaGram * 5
	} else {
		biayaTambahan = sisaGram * 15
	}

	totalBiaya := (beratKg * biayaPerKg) + biayaTambahan
	return beratKg, sisaGram, totalBiaya
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