package main

import (
	"fmt"
)

func main() {
	var berat int

	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scanf("%d", &berat)

	kg := berat / 1000
	grams := berat % 1000

	original := kg * 10000
	tambahan := 0

	if kg >= 10 {
		tambahan = 0
	} else {
		if grams < 500 {
			tambahan = grams * 15
		} else {
			tambahan = grams * 5
		}
	}

	total := original + tambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, grams)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", original, tambahan)
	fmt.Printf("Total biaya: Rp. %d\n", total)
}
