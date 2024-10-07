package main

import (
	"fmt"
	"strings"
)

func main() {
	urutanBenar := [4]string{"merah", "kuning", "hijau", "ungu"}
	berhasil := true

	for i := 0; i < 5; i++ {
		fmt.Printf("Percobaan %d: ", i+1)
		for j := 0; j < 4; j++ {
			var warna string
			fmt.Scan(&warna)
			if strings.ToLower(warna) != urutanBenar[j] {
				berhasil = false
			}
		}
	}

	fmt.Printf("BERHASIL: %t\n", berhasil)
}
