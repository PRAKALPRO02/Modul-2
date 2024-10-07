package main

import (
	"fmt"
	"strings"
)

func main() {
	var percobaan = true
	var nampung [4]string
	var warna = []string{"merah", "kuning", "hijau", "ungu"}
	for i := 0; i < 5; i++ {
		fmt.Printf("Percobaan ke-%v = ", i+1)
		for j := 0; j < 4; j++ {
			fmt.Scan(&nampung[j])
			strings.ToLower(nampung[j])
		}
		for j := 0; j < 4; j++ {
			if nampung[j] != warna[j] {
				percobaan = false
				break
			}
		}

	}
	fmt.Print("BERHASIL = ", percobaan)

}
