package main

import "fmt"

func main() {
	var beratGram int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratGram)

	// Konversi ke kilogram dan gram
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	// Hitung biaya dasar
	biayaDasar := beratKg * 10000

	// Hitung biaya tambahan
	var biayaTambahan int
	if sisaGram >= 500 {
		biayaTambahan = sisaGram * 5
	} else {
		biayaTambahan = sisaGram * 15
	}

	// Pengecualian untuk sisa berat kurang dari 1kg jika total berat lebih dari 10kg
	if beratKg > 10 && sisaGram < 1000 {
		biayaTambahan = 0
	}

	// Hitung total biaya
	totalBiaya := biayaDasar + biayaTambahan

	// Tampilkan hasil
	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}

//Zahrina Antika Malahati_2311102109
