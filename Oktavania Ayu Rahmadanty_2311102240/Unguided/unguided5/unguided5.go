package main

import "fmt"

func main() {
	var beratParsel int

	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&beratParsel)

	kg := beratParsel / 1000
	sisaGram := beratParsel % 1000

	biayaDasar := kg * 10000

	var biayaTambahan int
	if sisaGram >= 500 {
		biayaTambahan = sisaGram * 5
	} else {
		biayaTambahan = sisaGram * 15
	}

	if kg > 10 && sisaGram < 500 {
		biayaTambahan = 0
	}

	totalBiaya := biayaDasar + biayaTambahan

	fmt.Printf("Total berat: %d kg %d gram\n", kg, sisaGram)
	fmt.Printf("Biaya dasar: Rp %d\n", biayaDasar)
	fmt.Printf("Biaya tambahan: Rp %d\n", biayaTambahan)
	fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}
