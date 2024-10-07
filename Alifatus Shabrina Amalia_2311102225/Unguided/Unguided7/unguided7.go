package main

import (
	"fmt"
)

func main() {
	var beratParsel int
	const biayaPerKg = 10000
	const tambahanBiaya500gr = 5
	const tambahanBiayaKurang500gr = 15

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratParsel)

	beratKg := beratParsel / 1000
	sisaGram := beratParsel % 1000

	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)

	totalBiaya := beratKg * biayaPerKg

	if beratKg > 10 {

		fmt.Printf("Detail biaya: Rp. %d + Rp. 0\n", totalBiaya)
		totalBiaya += 0
	} else if sisaGram >= 500 {

		tambahanBiaya := sisaGram * tambahanBiaya500gr
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", totalBiaya, tambahanBiaya)
		totalBiaya += tambahanBiaya
	} else {

		tambahanBiaya := sisaGram * tambahanBiayaKurang500gr
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", totalBiaya, tambahanBiaya)
		totalBiaya += tambahanBiaya
	}

	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
