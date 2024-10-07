package main

import "fmt"


func hitungBiaya(beratGram int) (int, int, int) {
	
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	biayaKg := beratKg * 10000

	var biayaSisa int
	if sisaGram >= 500 {
		biayaSisa = sisaGram * 5
	} else if sisaGram > 0 && beratKg <= 10 {
		biayaSisa = sisaGram * 15
	} else {
		biayaSisa = 0
	}

	return biayaKg, biayaSisa, beratKg
}

func main() {
	var beratGram int
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scanln(&beratGram)

	
	biayaKg, biayaSisa, beratKg := hitungBiaya(beratGram)

	fmt.Printf("Total berat: %d kg %d gram\n", beratKg, beratGram%1000)
	fmt.Printf("Biaya pengiriman: Rp. %d\n", biayaKg+biayaSisa)
	fmt.Printf("  - Biaya per kg: Rp. %d\n", biayaKg)
	fmt.Printf("  - Biaya sisa gram: Rp. %d\n", biayaSisa)
}