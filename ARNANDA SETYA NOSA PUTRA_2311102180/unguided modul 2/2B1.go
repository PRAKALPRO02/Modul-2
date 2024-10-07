package main

import (
	"fmt"
	"strings"
)

func main() {
	const urutan = "merahkuninghijauungu"
	var gelas [4]string
	var hasil bool = true

	for i := 0; i < 5; i++ {

		fmt.Printf("Percobaan %d: ", i+1)
		for j := 0; j < 4; j++ {
			fmt.Scan(&gelas[j])
		}

		if urutan != strings.ToLower(strings.Join(gelas[:], "")) {
			hasil = false
		}
	}

	fmt.Printf("Berhasil: %t\n", hasil)
}
