package main

import (
	"fmt"
)

func main() {
	var beratParsel, biayaPengiriman, biayaTambahan, tambahanBiaya int

	fmt.Print("Masukkan berat parsel dalam gram: ")
	fmt.Scan(&beratParsel)

	kg := beratParsel / 1000
	sisaBerat := beratParsel % 1000
	biayaPengiriman = kg * 10000

	if sisaBerat >= 500 && sisaBerat < 1000 {
		tambahanBiaya = sisaBerat * 5
	} else if sisaBerat < 500 && sisaBerat > 0 {
		tambahanBiaya = sisaBerat * 15
	}

	biayaTambahan = biayaPengiriman + tambahanBiaya
	if kg > 10 {
		fmt.Printf("Total biaya pengiriman: Rp %d\n", biayaPengiriman)
	} else {
		fmt.Printf("Total biaya pengiriman: Rp %d\n", biayaTambahan)
	}
}
