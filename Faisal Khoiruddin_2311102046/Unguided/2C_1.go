package main

import "fmt"

func main() {
	var beratParsel int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratParsel)

	berat_kg := beratParsel / 1000
	sisa_gram := beratParsel % 1000

	biayaPerKg := 10000

	var tambahanBiaya int
	if sisa_gram >= 500 {
		tambahanBiaya = sisa_gram * 5
	} else {
		tambahanBiaya = sisa_gram * 15
	}

	totalBiaya := (berat_kg * biayaPerKg) + tambahanBiaya

	if berat_kg > 10 {
		totalBiaya -= tambahanBiaya
	}

	detailBerat := fmt.Sprintf("%d kg + %d gr", berat_kg, sisa_gram)
	detailBiaya := fmt.Sprintf("Rp. %d + Rp. %d", berat_kg*biayaPerKg, tambahanBiaya)

	fmt.Printf("Detail berat: %s\n", detailBerat)
	fmt.Printf("Detail biaya: %s\n", detailBiaya)
	fmt.Printf("Total biaya pengiriman: Rp. %d\n", totalBiaya)
}
