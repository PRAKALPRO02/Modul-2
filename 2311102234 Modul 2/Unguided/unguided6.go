package main

import "fmt"

func main() {
	var warna [4]string
	var berhasil bool = true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&warna[0], &warna[1], &warna[2], &warna[3])

		if warna[0] != "merah" || warna[1] != "kuning" || warna[2] != "hijau" || warna[3] != "ungu" {
			berhasil = false
		}
	}

	fmt.Printf("BERHASIL: %t\n", berhasil)
}
