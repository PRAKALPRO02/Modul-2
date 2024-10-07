package main

import "fmt"

func main() {
	var berat, Kg, gram, BiayaDasar, tambahan, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	Kg = berat / 1000
	gram = berat % 1000

	BiayaDasar = Kg * 10000

	if Kg > 10 {
		tambahan = 0
	} else {
		if gram >= 500 {
			tambahan = gram * 5
		} else {
			tambahan = gram * 15
		}
	}

	total = BiayaDasar + tambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", Kg, gram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", BiayaDasar, tambahan)
	fmt.Printf("Total biaya : Rp %d\n", total)
}
