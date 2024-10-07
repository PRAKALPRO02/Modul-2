package main

import (
	"fmt"
	"strings"
)

func cekUrutan(input []string) bool {
	const urutanBenar = "merahkuninghijauungu"
	urutanInput := strings.ToLower(strings.Join(input, ""))
	return urutanInput == urutanBenar
}

func main() {
	var berhasil bool = true
	var inputWarna [4]string

	for percobaan := 1; percobaan <= 5; percobaan++ {
		fmt.Printf("Percobaan %d: ", percobaan)
		for i := 0; i < 4; i++ {
			fmt.Scan(&inputWarna[i])
		}

		if !cekUrutan(inputWarna[:]) {
			berhasil = false
		}
	}

	fmt.Printf("Berhasil: %t\n", berhasil)
}
