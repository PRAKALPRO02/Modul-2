package main

import (
	"fmt"
	"strings"
)

func main() {
	// Deklarasi urutan warna yang benar
	warnaBenar := [4]string{"merah", "kuning", "hijau", "ungu"}
	var berhasil bool = true

	// Loop untuk 5 kali percobaan
	for i := 1; i <= 5; i++ {
		var warnaInput [4]string
		fmt.Printf("Percobaan %d: ", i)
		for j := 0; j < 4; j++ {
			fmt.Scan(&warnaInput[j])
		}

		// Cek apakah input sama dengan urutan warna yang benar
		for k := 0; k < 4; k++ {
			if strings.ToLower(warnaInput[k]) != warnaBenar[k] {
				berhasil = false
			}
		}
	}

	// Output hasil
	if berhasil {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
