package main

import "fmt"

func main() {
	warnaIdeal := []string{"merah", "kuning", "hijau", "ungu"}
	berhasil := 0

	for i := 1; i < 6; i++ {
		fmt.Print("Percobaan ", i, " : ")
		var warna [4]string
		fmt.Scan(&warna[0], &warna[1], &warna[2], &warna[3])

		isSesuai := true
		for j := 0; j < len(warna); j++ {
			if warna[j] != warnaIdeal[j] {
				isSesuai = false
				break
			}
		}
		if isSesuai {
			berhasil++
		}
	}

	if berhasil == 5 {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
