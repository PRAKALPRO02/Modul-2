package main

import (
	"fmt"
)

func main() {
	var beratParsel int
	var beratKg, sisaBerat int
	var biayaKg, biayaSisa, totalBiaya int

	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scanln(&beratParsel)

	beratKg = beratParsel / 1000
	sisaBerat = beratParsel % 1000

	biayaKg = beratKg * 10000

	if beratKg > 10 {
		biayaSisa = 0
	} else if sisaBerat >= 500 {
		biayaSisa = sisaBerat * 5
	} else {
		biayaSisa = sisaBerat * 15
	}

	totalBiaya = biayaKg + biayaSisa

	fmt.Println("Detail berat:", beratKg, "kg +", sisaBerat, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp.", totalBiaya)
}
