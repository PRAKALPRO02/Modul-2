package main

import "fmt"

func main() {
	var beratParsel, kg, sisaGram, biayaDasar, biayaTambahan , totalBiaya int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratParsel)

	kg = beratParsel / 1000
	sisaGram = beratParsel % 1000

	biayaDasar = kg * 10000

	if sisaGram >= 500 {
			biayaTambahan = sisaGram * 5
	} else if sisaGram > 0 {
			biayaTambahan = sisaGram * 15
	}

	if kg > 10 && sisaGram < 1000 {
			biayaTambahan = 0
	}
	
	totalBiaya = biayaDasar + biayaTambahan
	fmt.Println("Detail berat :",kg,"kg +", sisaGram, "gr")
	fmt.Println("Detail biaya : Rp.",biayaDasar, "+ Rp.", biayaTambahan)
	fmt.Println("Total biaya : Rp.",totalBiaya)
}
