package main

import "fmt"

func main() {
	var berat int
	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	gram := berat % 1000

	biayaTotal := kg * 10000
	biayaTambahan := 0

	if gram >= 500 {
		biayaTambahan = gram * 5
	} else if gram > 0 {
		biayaTambahan = gram * 15
	}

	if kg >= 10 {
		biayaTambahan = 0
	}

	biayaTotal += biayaTambahan

	fmt.Printf("Detail berat: %d kg + %d gram\n", kg, gram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", kg*10000, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", biayaTotal)
}
