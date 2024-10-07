package main

import (
	"fmt"
	"strings"
)

func main() {
	var percobaan = true
	var E [4]string
	for i := 0; i < 5; i++ {
		fmt.Printf("Percobaan ke-%v = ", i+1)
		for j := 0; j < 4; j++ {
			fmt.Scan(&E[j])
			strings.ToLower(E[j])
		}
		if E[0] != "merah" && E[1] != "kuning" && E[2] != "hijau" && E[3] != "ungu" {
			percobaan = false
		}

	}
	fmt.Print("BERHASIL = ", percobaan)

}
