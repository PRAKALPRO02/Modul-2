package main

import "fmt"

func main() {
	var beratParsel, parselKG, parselGram, biayaKg, biayaGram int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratParsel)
	parselKG = beratParsel / 1000
	parselGram = beratParsel % 1000
	biayaKg = parselKG * 10000
	if parselGram >= 15 {
		biayaGram = parselGram * 5
	} else {
		biayaGram = parselGram * 15
	}
	fmt.Printf("Detail berat: %d kg + %d gr\n", parselKG, parselGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaGram)
	fmt.Printf("Total biaya: Rp. %d\n", biayaKg+biayaGram)
}
