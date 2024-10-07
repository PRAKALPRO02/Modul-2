package main

import "fmt"

func main() {
	var beratParsel int
	var beratKg, sisaGram, biayaKg, biayaTambahan, totalBiaya int

	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scan(&beratParsel)

	beratKg = beratParsel / 1000
	sisaGram = beratParsel % 1000

	biayaKg = beratKg * 10000

	// Kondisi untuk menentukan biaya tambahan
	if sisaGram > 0 {
		if sisaGram >= 500 {
			biayaTambahan = sisaGram * 5
		} else {
			biayaTambahan = sisaGram * 15
		}
	}

	totalBiaya = biayaKg + biayaTambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
