package main

import (
	"fmt"
)

func hitungBiayaPos(berat int) (int, string, int) {

	kg := berat / 1000
	sisaGram := berat % 1000

	biayaPerKg := 10000
	biayaTambahan := 0

	totalBiaya := kg * biayaPerKg

	if sisaGram > 500 {
		biayaTambahan = sisaGram * 15 // Jika lebih dari 500 gram, biaya tambahan Rp. 15,- per gram
	} else if sisaGram > 0 {
		biayaTambahan = sisaGram * 5 // Jika kurang dari 500 gram, biaya tambahan Rp. 5,- per gram
	}

	if kg >= 10 {
		biayaTambahan = 0
	}

	totalBiaya += biayaTambahan

	detailBerat := fmt.Sprintf("%d kg + %d gr", kg, sisaGram)

	return totalBiaya, detailBerat, biayaTambahan
}

func main() {
	var berat int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	totalBiaya, detailBerat, biayaTambahan := hitungBiayaPos(berat)

	fmt.Println("Detail berat:", detailBerat)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", (totalBiaya - biayaTambahan), biayaTambahan)
	fmt.Println("Total biaya: Rp.", totalBiaya)
}
