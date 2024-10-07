package main

import "fmt"

var (
	berat_Parsel, beratKg, sisaGram, biayaKg, biayaTambahan, total_Biaya int
)

func main() {

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat_Parsel)

	beratKg = berat_Parsel / 1000
	sisaGram = berat_Parsel % 1000

	biayaKg = beratKg * 10000

	if beratKg > 10 {
		biayaTambahan = 0
	} else {
		if sisaGram >= 500 {
			biayaTambahan = sisaGram * 5
		} else {
			biayaTambahan = sisaGram * 15
		}
	}

	total_Biaya = biayaKg + biayaTambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", total_Biaya)
}
