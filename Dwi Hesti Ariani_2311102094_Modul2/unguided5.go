package main

import "fmt"

func main() {
	var beratParsel int
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scanln(&beratParsel)

	totalBeratKg := beratParsel / 1000
	sisaBeratGram := beratParsel % 1000

	biayaPerKg := 10000
	biayaTotalKg := totalBeratKg * biayaPerKg

	if sisaBeratGram >= 500 {
		biayaSisaGram := sisaBeratGram * 5
		biayaTotal := biayaTotalKg + biayaSisaGram
		fmt.Println("Biaya total:", biayaTotal)
	} else if sisaBeratGram < 500 {
		if totalBeratKg > 10 {
			biayaTotal := biayaTotalKg
			fmt.Println("Biaya total:", biayaTotal)
		} else {
			biayaSisaGram := sisaBeratGram * 15
			biayaTotal := biayaTotalKg + biayaSisaGram
			fmt.Println("Biaya total:", biayaTotal)
		}
	}
}

//Dwi Hesti Ariani_2311102094
