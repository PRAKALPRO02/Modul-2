package main

import (
	"fmt"
)

func main() {
	var berat, biaya, val, tambahan int

	fmt.Print("Masukkan berat parsel dalam gram: ")
	fmt.Scan(&berat)

	kg := berat / 1000
	sisaBerat := berat % 1000
	biaya = kg * 10000

	if sisaBerat >= 500 && sisaBerat < 1000 {
		for i := 0; i < sisaBerat; i++ {
			tambahan = tambahan + 5
		}
	} else if sisaBerat < 500 && sisaBerat > 0 {
		for i := 0; i < sisaBerat; i++ {
			tambahan = tambahan + 15
		}
	}

	val = biaya + tambahan
	if kg > 10 {

		fmt.Printf("Total biaya pengiriman: Rp %d\n", biaya)
	} else {
		fmt.Printf("Total biaya pengiriman: Rp %d\n", val)
	}

}
