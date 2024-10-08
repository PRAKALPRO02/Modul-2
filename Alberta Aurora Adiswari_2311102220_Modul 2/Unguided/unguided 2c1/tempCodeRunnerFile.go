package main

import (
	"fmt"
)

func main() {
	var beratParcel int

	// Meminta input berat parsel dalam gram
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratParcel)

	// Menghitung total berat dalam kg dan sisa dalam gram
	totalKg_220 := beratParcel / 1000
	sisaGr_220 := beratParcel % 1000

	// Menghitung biaya dasar
	biayaDasar := totalKg_220 * 10000
	var biayaSisa int

	// Menghitung biaya tambahan berdasarkan sisa berat
	if totalKg_220 > 10 {
		biayaSisa = 0 // Sisa berat gratis jika total lebih dari 10kg
	} else {
		if sisaGr_220 >= 500 {
			biayaSisa = sisaGr_220 * 5 // Rp. 5,- per gram
		} else {
			biayaSisa = sisaGr_220 * 15 // Rp. 15,- per gram
		}
	}

	// Menampilkan detail berat dan biaya
	fmt.Printf("Detail berat: %d kg %d gr\n", totalKg_220, sisaGr_220)
	fmt.Printf("Detail biaya: Rp. %d Rp. %d\n", biayaDasar, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", biayaDasar+biayaSisa)
}
